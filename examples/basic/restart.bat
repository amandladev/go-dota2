@echo off
:restart
echo Reiniciando o aplicativo Go...

:: Encerrar o processo atual se ele já estiver em execução
:: Você pode precisar ajustar isso se o nome do processo não for 'main.exe'
taskkill /IM main.exe /F
taskkill /FI "WINDOWTITLE eq C:\Users\Servidor\Desktop\asd\go-dota2\examples\basic - main.go*" /F

:: Navegar até o diretório do seu aplicativo
cd C:\Users\Servidor\Desktop\asd\go-dota2\examples\basic

:: Iniciar o aplicativo
start "C:\Users\Servidor\Desktop\asd\go-dota2\examples\basic - main.go" cmd /c go run main.go

:: Aguardar 3 horas (10800 segundos) antes de reiniciar
timeout /t 3600

goto restart
