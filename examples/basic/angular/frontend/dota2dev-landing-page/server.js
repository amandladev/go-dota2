const express = require('express');
const path = require('path');
const app = express();

const PORT = 30125;

// Servir os arquivos estáticos do aplicativo Angular
app.use(express.static(path.join(__dirname, 'dist/dota2dev-landing-page/browser')));

// Rota para qualquer outra requisição que não seja arquivos estáticos
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'dist/dota2dev-landing-page/browser/index.html'));
});

// Iniciar o servidor na porta 30125
app.listen(PORT, () => {
  console.log(`Servidor rodando na porta ${PORT}`);
});
