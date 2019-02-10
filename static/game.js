const startGame = document.getElementById('game-start');

const declareWinner = document.getElementById('declare-winner');
const submitWinnerButton = document.getElementById('winner-button');
const winnerInput = document.getElementById('winner');

const blindContainer = document.getElementById('blind-value');

const gameContainer = document.getElementById('game');
const gameEndContainer = document.getElementById('game-end');

declareWinner.hidden = true;
gameEndContainer.hidden = true;

const synth = window.speechSynthesis;

document.getElementById('start-game').addEventListener('click', event => {
    startGame.hidden = true;
    declareWinner.hidden = false;

    const utterThis = new SpeechSynthesisUtterance('Lets play poker');
    synth.speak(utterThis);

    const numberOfPlayers = document.getElementById('player-count').value;

    if (window['WebSocket']) {
        const conn = new WebSocket('ws://' + document.location.host + '/ws');

        submitWinnerButton.onclick = event => {
            conn.send(winnerInput.value);
            synth.speak(new SpeechSynthesisUtterance(`Congratulations ${winnerInput.value}`));
            gameEndContainer.hidden = false;
            gameContainer.hidden = true
        };

        conn.onclose = evt => {
            blindContainer.innerText = 'Connection closed'
        };

        conn.onmessage = evt => {
            const data = evt.data;
            synth.speak(new SpeechSynthesisUtterance(data));
            blindContainer.innerText = data
        };

        conn.onopen = () => {
            conn.send(numberOfPlayers)
        }
    }
});
