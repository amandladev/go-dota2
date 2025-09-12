package main

import (
	//"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync" // Importante para usar o mutex
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/paralin/go-dota2"
	//"github.com/paralin/go-dota2/cso"
	"github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-steam"
	"github.com/paralin/go-steam/protocol/steamlang"
	steamId "github.com/paralin/go-steam/steamid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

const steamID64Offset uint64 = 76561197960265728


type handler struct {
	steamClient *steam.Client
	dotaClient  *dota2.Dota2
	mu          sync.Mutex // Mutex para controlar o acesso exclusivo
	//counter int // Contador global para gerar identificadores únicos para o contexto
}

func main() {
    loadEnv()

    // Initialize the handler with empty clients
    handler := handler{}

	// Depois, inicie o servidor Gin
    initGinServer(&handler) // O servidor será iniciado aqui e bloqueará o main

    // Inicie a conexão com a Steam primeiro
    initSteamConnection(&handler)

	// Bloqueia a execução para evitar que o programa termine
    select {}
}

func errorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			log.Printf("Errors found in request: %v", c.Errors)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}
}

func securityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if suspicious(c.Request.URL.Path) {
			log.Printf("Blocked suspicious request: %v", c.Request.URL.Path)
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}

func suspicious(uri string) bool {
	return strings.Contains(uri, "phpmyadmin") ||
		strings.Contains(uri, "cgi-bin") ||
		strings.Contains(uri, "config")
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	})
}

func initSteamConnection(h *handler) {
    // Set Steam credentials from the environment
    details := &steam.LogOnDetails{
        Username: os.Getenv("BOT_USERNAME"),
        Password: os.Getenv("BOT_PASSWORD"),
    }

    // Initialize Steam directory
    err := steam.InitializeSteamDirectory()
    if err != nil {
        panic(err)
    }

    // Initialize the Steam client
    h.steamClient = steam.NewClient()
    h.steamClient.Connect()

    // Start the event loop
    h.startEventLoop(details)
}

func (h *handler) startEventLoop(details *steam.LogOnDetails) {
    go func() {
        for event := range h.steamClient.Events() {
            switch e := event.(type) {
            case *steam.ConnectedEvent:
                //log.Println("Connected to Steam network, logging in...")
                h.steamClient.Auth.LogOn(details)
            case *steam.LoggedOnEvent:
                //log.Println("Successfully logged on to Steam")

                // Set account state to online
                h.steamClient.Social.SetPersonaState(steamlang.EPersonaState_Online)

                // Initialize the Dota 2 client
                h.dotaClient = dota2.New(h.steamClient, logrus.New())
                h.dotaClient.SetPlaying(true)

                //log.Println("Dota 2 client initialized")
                h.dotaClient.SayHello()
            case steam.FatalErrorEvent:
                //log.Println("Fatal error occurred: ", e)
                // Handle reconnection logic if needed
            case error:
                log.Println("Error: ", e)
            }
        }
    }()
}


// Função para reinicializar o estado da party
func (h *handler) resetPartyState() {
    //log.Println("Reinicializando o estado da party...")
    h.dotaClient.LeaveParty()
    time.Sleep(2 * time.Second) // Aguarda para garantir que a party foi completamente destruída
    //log.Println("Estado da party reinicializado.")
}

func (h *handler) resetSteamConnection() {
    //log.Println("Reiniciando a conexão com a Steam...")
    h.steamClient.Disconnect()
    time.Sleep(2 * time.Second) // Aguarda para garantir a desconexão

    // Reconecta à Steam
    initSteamConnection(h)
    //log.Println("Conexão com a Steam reiniciada com sucesso")
}

// Função para analisar o Steam ID e converter para 64 bits se necessário
func parseSteamID(idStr string) (uint64, error) {
    steamID64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        return 0, err
    }

    // Se o SteamID for menor que o offset, é um ID de 32 bits e precisa de conversão
    if steamID64 < steamID64Offset {
        //log.Printf("Recebido Steam ID de 32 bits, convertendo para 64 bits")
        steamID64 += steamID64Offset
    }
    return steamID64, nil
}

// Função opcional: Converter SteamID64 para AccountID (32 bits)
func convertSteamID64ToAccountID(steamID64 uint64) uint32 {
    return uint32(steamID64 - steamID64Offset)
}

// Função principal para processar as ações do grupo
func (h *handler) processGroupActions(steamIDs []string) error {
    //log.Println("Iniciando o processamento das ações do grupo...")

    // Reinicializa o estado da party antes de começar
    h.resetPartyState()
    time.Sleep(2 * time.Second)

    // Convida os jogadores para a party
    for _, idStr := range steamIDs {
        steamID64, err := parseSteamID(idStr)
        if err != nil {
            log.Printf("Steam ID inválido: %s", idStr)
            continue
        }

        // Convida o jogador para a party
        h.dotaClient.InviteToParty(steamID64)
        steamID32 := convertSteamID64ToAccountID(steamID64)
    	log.Printf("Steam ID 32: %d, Steam ID 64: %d", steamID32, steamID64)

    }

    // Aguarda 30 segundos para que os jogadores entrem
    //log.Println("Aguardando 30 segundos para os jogadores entrarem...")
    time.Sleep(30 * time.Second)//padrao 30
	/*PROBLEMA INICIA AKI
    // Realiza o ready check com timeout aumentado
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)//padrao 30
    defer cancel()

    readyCheckResp, err := h.dotaClient.SendPartyReadyCheck(ctx)
    if err != nil {
        log.Printf("Erro ao iniciar o ready check: %v", err)
        return err
    }
    log.Printf("Ready check iniciado: %v", readyCheckResp)
	*///PROBLEMA ACABA AKI

    // Define o líder da party como o primeiro Steam ID
    leaderID64, err := parseSteamID(steamIDs[0])
    if err != nil {
        //log.Printf("Steam ID do líder inválido: %s", steamIDs[0])
        return err
    }

    h.dotaClient.SetPartyLeader(steamId.SteamId(leaderID64))
	leaderID32 := convertSteamID64ToAccountID(leaderID64)
	log.Printf("Líder Steam ID 32 bits: %d, Steam ID 64 bits: %d", leaderID32, leaderID64)

    // Aguarda um pouco para garantir que o líder foi definido
    time.Sleep(1 * time.Second)

    // Sai da party para resetar o estado
    h.dotaClient.LeaveParty()
    //log.Println("Saiu da party")

    time.Sleep(2 * time.Second)

    // Reinicializa a conexão com a Steam para evitar inconsistências
    h.resetSteamConnection()

    //log.Println("Processamento das ações do grupo concluído com sucesso")

    // Não enviar resposta HTTP aqui

    return nil
}

func initGinServer(handler *handler) {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(errorHandlingMiddleware()) // Aqui o middleware é aplicado ao router
	r.Use(securityMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	})

	// Create a simple index.html page
	r.LoadHTMLGlob("templates/*")

	r.Static("/images", "/app/examples/basic/images")
	r.Static("/static", "/app/examples/basic/angular/frontend/test/dist/test/browser/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.File("/app/examples/basic/angular/frontend/test/dist/test/browser/index.html")
	})
	/**
	Create Lobby
	Check types to custom your lobby -> protocol/dota_gcmessages_client_match_management.pb.go > CMsgPracticeLobbySetDetails
	*/
	r.POST("/lobby", func(c *gin.Context) {

		lobbyVisibility := protocol.DOTALobbyVisibility_DOTALobbyVisibility_Public

		// Peruvian Server ID is # 15, choose the server id you prefer...
		lobbyRegion := uint32(10)

		// Captains Mode is # 2
		gameMode := uint32(2)

		lobbyDetails := &protocol.CMsgPracticeLobbySetDetails{
			GameName:     proto.String("www.kronon.com.br"),
			Visibility:   &lobbyVisibility,
			PassKey:      proto.String("kronon"),
			ServerRegion: &lobbyRegion,
			GameMode:     &gameMode,
		}

		handler.dotaClient.CreateLobby(lobbyDetails)

		uintId, err := strconv.ParseUint(os.Getenv("BOT_STEAMID"), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		steamId := uint32(uintId)

		// For some reason the bot automatically joins the first slot. Kick him.
		handler.dotaClient.KickLobbyMemberFromTeam(steamId)

		c.JSON(200, gin.H{
			"message": "Lobby has been created",
		})
	})

	// Invite a player to the lobby
	r.POST("/invite/:steamId", func(c *gin.Context) {
		// Extrai o steamId da URL.
		steamIdStr := c.Param("steamId")
		log.Printf("Inviting player with steamId: %v", steamIdStr)

		// Converte o steamIdStr para um uint64.
		steamId64, err := strconv.ParseUint(steamIdStr, 10, 64)
		if err != nil {
			log.Printf("Error converting steamId to uint64: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SteamID format"})
			return
		}

		// Assumindo que você tem uma função adequada para converter uint64 para steamid.SteamId
		// Isso depende da implementação específica do seu tipo steamid.SteamId.
		sid := steamId.SteamId(steamId64) // Este é um exemplo; ajuste conforme a sua implementação específica.

		// Chama InviteLobbyMember com o sid correto.
		handler.dotaClient.InviteLobbyMember(sid)

		c.JSON(http.StatusOK, gin.H{
			"message": "Player has been invited",
		})
	})

	// Start game in current lobby
	r.POST("/start", func(c *gin.Context) {
		handler.dotaClient.LaunchLobby()
		c.JSON(200, gin.H{
			"message": "Lobby has been started",
		})
	})

	// Leave current lobby
	r.POST("/leave", func(c *gin.Context) {
		handler.dotaClient.LeaveLobby()
		c.JSON(200, gin.H{
			"message": "You have left the room",
		})
	})

	r.POST("/chamada-unica", func(c *gin.Context) {
		// Tenta adquirir o lock (mutex)
		handler.mu.Lock()
		defer handler.mu.Unlock()
	
		var request struct {
			SteamIDs []string `json:"steamIds"`
		}
	
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de requisição inválido"})
			return
		}
	
		if len(request.SteamIDs) < 2 || len(request.SteamIDs) > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "É necessário fornecer entre 2 e 5 Steam IDs"})
			return
		}
	
		// Processa as ações com os Steam IDs fornecidos
		err := handler.processGroupActions(request.SteamIDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		// Envia a resposta ao cliente aqui
		c.JSON(http.StatusOK, gin.H{
			"message": "Processo concluído com sucesso",
			"detalhes": "www.kronon.com.br",
		})
	})	

	// Start the web server
	go func() {
		err := r.Run(":8080")
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}