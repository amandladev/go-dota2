document.getElementById('createRoom').addEventListener('click', createRoom);
document.getElementById('deleteAllRooms').addEventListener('click', deleteAllRooms);

function fetchRooms() {
    // Simulação de fetch, substitua com sua API
    fetch('/api/rooms')
        .then(response => response.json())
        .then(rooms => {
            const roomsContainer = document.getElementById('roomsContainer');
            roomsContainer.innerHTML = '';
            rooms.forEach(room => {
                const roomElement = document.createElement('div');
                roomElement.className = 'room';
                roomElement.innerHTML = `
                    <h3>${room.name}</h3>
                    <button onclick="deleteRoom('${room.id}')">Deletar Sala</button>
                `;
                roomsContainer.appendChild(roomElement);
            });
        });
}

function createRoom() {
    // Simulação de post, substitua com sua API
    const roomName = prompt("Nome da Sala:");
    fetch('/api/rooms', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ name: roomName })
    }).then(() => fetchRooms());
}

function deleteAllRooms() {
    // Simulação de delete, substitua com sua API
    fetch('/api/rooms', { method: 'DELETE' }).then(() => fetchRooms());
}

function deleteRoom(id) {
    // Simulação de delete, substitua com sua API
    fetch(`/api/rooms/${id}`, { method: 'DELETE' }).then(() => fetchRooms());
}

fetchRooms(); // Carrega inicialmente todas as salas
