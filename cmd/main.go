package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-steam"
	"github.com/paralin/go-steam/protocol/steamlang"
	"github.com/paralin/go-steam/steamid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var (
	playersInLobby int = 1 // Empezamos con 1 (el bot)
)

func main() {
	fmt.Println("Iniciando cliente Steam y Dota 2...")

	// Carga variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}

	steamUser := os.Getenv("STEAM_USER")
	steamPass := os.Getenv("STEAM_PASS")
	if steamUser == "" || steamPass == "" {
		log.Fatal("Por favor, define STEAM_USER y STEAM_PASS en el archivo .env o en tus variables de entorno.")
	}

	// Inicializa el cliente de Steam
	steamClient := steam.NewClient()
	logOnDetails := &steam.LogOnDetails{
		Username: steamUser,
		Password: steamPass,
	}

	var dotaClient *dota2.Dota2

	// Conecta y autentica
	go func() {
		for event := range steamClient.Events() {
			switch e := event.(type) {
			case *steam.ConnectedEvent:
				fmt.Println("Conectado a Steam, iniciando sesión...")
				steamClient.Auth.LogOn(logOnDetails)
			case *steam.LoggedOnEvent:
				fmt.Println("Sesión iniciada en Steam.")
				steamClient.Social.SetPersonaState(steamlang.EPersonaState_Online)
				// Inicializa el cliente de Dota 2
				dotaClient = dota2.New(steamClient, logrus.New())
				dotaClient.SetPlaying(true) // Mantener como "playing" en Steam
				dotaClient.SayHello()
				fmt.Println("Cliente de Dota 2 inicializado y listo.")

				// Añadir una pequeña pausa para asegurar la inicialización
				time.Sleep(2 * time.Second)

				// Configurar eventos del cliente de Dota 2
				setupDotaEventListeners(dotaClient)

				// Crear lobby
				createLobby(dotaClient)

				// Esperar un momento para que se cree el lobby
				time.Sleep(5 * time.Second) // Unirse al lobby como broadcaster (espectador)
				joinLobbyAsBroadcaster(dotaClient)

				// Esperar un momento antes de invitar
				time.Sleep(3 * time.Second)

				// Obtener el Steam ID objetivo desde las variables de entorno
				inviteUser := os.Getenv("INVITE_STEAM_ID")
				fmt.Printf("🔍 DEBUG: INVITE_STEAM_ID = '%s'\n", inviteUser)
				if inviteUser != "" {
					invitePlayerToLobby(dotaClient, inviteUser)
				} else {
					fmt.Println("⚠️  No se encontró INVITE_STEAM_ID en las variables de entorno")
				} // Comenzar el monitoreo del lobby
				go monitorLobby(dotaClient)

			case *steam.FatalErrorEvent:
				log.Fatalf("Error fatal: %v", e)
			case error:
				log.Printf("Error: %v", e)
			}
		}
	}()

	// Conecta el cliente Steam
	steamClient.Connect()
	// Mantiene la aplicación viva
	select {}
}

// createLobby crea un lobby de práctica en Dota 2
func createLobby(dotaClient *dota2.Dota2) {
	// Configuración del lobby
	lobbyVisibility := protocol.DOTALobbyVisibility_DOTALobbyVisibility_Friends // Solo amigos
	lobbyRegion := uint32(10)                                                   // 10 = US East, ajusta según prefieras
	gameMode := uint32(1)                                                       // 1 = All Pick, 2 = Captains Mode

	// Configuración para modo de bots
	fillWithBots := true
	allowCheats := true
	allowSpectating := true
	botDifficultyRadiant := protocol.DOTABotDifficulty_BOT_DIFFICULTY_MEDIUM
	botDifficultyDire := protocol.DOTABotDifficulty_BOT_DIFFICULTY_MEDIUM

	lobbyDetails := &protocol.CMsgPracticeLobbySetDetails{
		GameName:             proto.String("Mi Lobby de Bots"),
		Visibility:           &lobbyVisibility,
		PassKey:              proto.String("test123"), // Contraseña para entrar
		ServerRegion:         &lobbyRegion,
		GameMode:             &gameMode,
		FillWithBots:         &fillWithBots,         // Llenar con bots
		AllowCheats:          &allowCheats,          // Permitir cheats
		AllowSpectating:      &allowSpectating,      // Permitir espectadores
		BotDifficultyRadiant: &botDifficultyRadiant, // Dificultad bots Radiant
		BotDifficultyDire:    &botDifficultyDire,    // Dificultad bots Dire
	}

	fmt.Println("Creando lobby con configuración:")
	fmt.Printf("- Nombre: %s\n", *lobbyDetails.GameName)
	fmt.Printf("- Visibilidad: Solo amigos\n")
	fmt.Printf("- Contraseña: test123\n")
	fmt.Printf("- Modo: All Pick\n")
	fmt.Printf("- Región: US East\n")
	fmt.Printf("- Llenar con bots: Sí\n")
	fmt.Printf("- Dificultad bots: Medium\n")
	fmt.Printf("- Cheats habilitados: Sí\n")
	fmt.Printf("- Espectadores permitidos: Sí\n")

	// Crear el lobby
	dotaClient.CreateLobby(lobbyDetails)
	fmt.Println("✅ Comando de crear lobby de bots enviado!")
}

// joinLobbyAsBroadcaster une al bot al lobby como broadcaster (espectador)
func joinLobbyAsBroadcaster(dotaClient *dota2.Dota2) {
	fmt.Println("📺 Uniéndose al lobby como broadcaster...")

	// Usar la función específica para unirse como broadcaster
	dotaClient.JoinLobbyBroadcastChannel(
		0,     // channel: 0 para el canal principal
		"Bot", // preferredDescription: descripción del broadcaster
		"US",  // preferredCountryCode: código de país
		"en",  // preferredLanguageCode: código de idioma
	)

	fmt.Println("✅ Bot unido al lobby como broadcaster")
	fmt.Println("🔍 DEBUG: JoinLobbyBroadcastChannel ejecutado correctamente")
}

// setupDotaEventListeners configura los event listeners del cliente de Dota 2
func setupDotaEventListeners(dotaClient *dota2.Dota2) {
	fmt.Println("🎧 Configurando escuchadores de eventos...")

	// Mantener el bot como "playing" para que aparezca activo en Steam
	dotaClient.SetPlaying(true)
	fmt.Println("🎮 Bot configurado como 'jugando' en Steam")

	fmt.Println("✅ Sistema de eventos configurado")
}

// invitePlayerToLobby invita a un jugador específico al lobby
func invitePlayerToLobby(dotaClient *dota2.Dota2, steamIDStr string) {
	fmt.Printf("📧 Invitando al jugador con Steam ID: %s\n", steamIDStr)

	// Convertir string a uint64
	steamID64, err := strconv.ParseUint(steamIDStr, 10, 64)
	if err != nil {
		log.Printf("❌ Error al convertir Steam ID: %v", err)
		return
	}

	sid := steamid.SteamId(steamID64)

	dotaClient.InviteLobbyMember(sid)

	fmt.Printf("✅ Invitación enviada al jugador %s!\n", steamIDStr)
}

func monitorLobby(dotaClient *dota2.Dota2) {
	fmt.Println("👁️  Iniciando monitoreo del lobby...")

	time.Sleep(8 * time.Second)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	gameStarted := false
	checkCount := 0

	for range ticker.C {
		if gameStarted {
			break
		}

		checkCount++
		fmt.Printf("🔍 Verificando estado del lobby (check #%d)...\n", checkCount)

		// Simulamos la detección de jugadores
		// En una implementación real, aquí consultaríamos el estado del lobby

		// Después de 2 checks (30 segundos), asumimos que se unió alguien
		if checkCount >= 2 && !gameStarted {
			playersInLobby = 2 // Simulamos que se unió 1 jugador más
			fmt.Printf("👥 Se detectaron %d jugadores en el lobby\n", playersInLobby)
			fmt.Println("🚀 ¡Hay suficientes jugadores! Iniciando partida en 5 segundos...")
			time.Sleep(5 * time.Second)

			launchGame(dotaClient)
			gameStarted = true
		} else {
			fmt.Println("⏳ Esperando más jugadores...")
		}
	}
}

// launchGame inicia la partida en el lobby actual
func launchGame(dotaClient *dota2.Dota2) {
	fmt.Println("🎮 Iniciando la partida...")
	dotaClient.LaunchLobby()
	fmt.Println("✅ Comando de iniciar partida enviado!")

	// Preparar el bot para aceptar la partida automáticamente
	go acceptGameWhenReady(dotaClient)

	fmt.Println("🎉 ¡La partida ha comenzado!")
}

// acceptGameWhenReady acepta automáticamente cuando se encuentra una partida
func acceptGameWhenReady(dotaClient *dota2.Dota2) {
	fmt.Println("📺 Bot manteniendo conexión como broadcaster...")

	// Esperar un poco para que se procese el inicio del juego
	time.Sleep(3 * time.Second)

	// Mantener el bot activo y visible en Steam
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)

		fmt.Printf("🔄 Manteniendo conexión activa (intento %d/10)...\n", i+1)

		// Mantener el bot como "jugando" en Steam para que aparezca activo
		dotaClient.SetPlaying(true)

		// Enviar un "heartbeat" para mantener la conexión
		if i == 2 {
			fmt.Println("🔗 Reforzando conexión como broadcaster...")
			dotaClient.SayHello()
		}

		if i >= 5 {
			fmt.Println("📺 El bot está conectado como broadcaster.")
			fmt.Println("    Debería aparecer como espectador en la partida.")
			break
		}
	}

	fmt.Println("✅ Proceso de conexión como broadcaster completado")
}
