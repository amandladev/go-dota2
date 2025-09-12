document.addEventListener('DOMContentLoaded', function() {
    document.querySelector('button').addEventListener('click', fetchMatchDetails);

    function fetchMatchDetails() {
        var matchId = document.getElementById('matchIdInput').value;
        if (!matchId) {
            alert('Por favor, insira um ID de partida');
            return;
        }
        fetch(`https://api.opendota.com/api/matches/${matchId}`)
            .then(response => response.json())
            .then(data => displayData(data))
            .catch(error => {
                console.error('Erro ao buscar dados:', error);
                document.getElementById('result').innerHTML = 'Falha ao carregar dados';
            });
    }

    function displayData(data) {
        const resultContainer = document.getElementById('result');
        resultContainer.innerHTML = '<h2><i class="fas fa-info-circle icon"></i>Detalhes da Partida</h2>';
        // Include additional game data...
    }
});

function createRoom() {
    // Implemente a lógica para criar uma sala
}

function deleteAllRooms() {
    // Implemente a lógica para deletar todas as salas
}
