const startGame = document.getElementById('game-start');

const declareWinner = document.getElementById('declare-winner');
const submitWinnerButton = document.getElementById('winner-button');
const winnerInput = document.getElementById('winner');

const blindContainer = document.getElementById('blind-value');

const gameContainer = document.getElementById('game');
const gameEndContainer = document.getElementById('game-end');

declareWinner.hidden = true;
gameEndContainer.hidden = true;

document.getElementById('start-game').addEventListener('click', event => {
    startGame.hidden = true;
    declareWinner.hidden = false;

    say('Lets play poker');

    const numberOfPlayers = document.getElementById('player-count').value;

    if (window['WebSocket']) {
        const conn = new WebSocket('ws://' + document.location.host + '/ws');

        submitWinnerButton.onclick = event => {
            conn.send(winnerInput.value);
            say(`Congratulations ${winnerInput.value}`);
            gameEndContainer.hidden = false;
            gameContainer.hidden = true
        };

        conn.onclose = evt => {
            blindContainer.innerText = 'Connection closed'
        };

        conn.onmessage = evt => {
            const data = evt.data;
            say(data);
            blindContainer.innerText = data
        };

        conn.onopen = () => {
            conn.send(numberOfPlayers)
        }
    }
});

function say(msg) {
    window.speechSynthesis.speak(new SpeechSynthesisUtterance(msg));
}