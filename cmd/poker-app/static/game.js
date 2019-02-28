const startGame = document.getElementById('game-start');

const declareWinner = document.getElementById('declare-winner');
const submitWinnerButton = document.getElementById('winner-button');
const winnerInput = document.getElementById('winner');

const blindContainer = document.getElementById('blind-value');

const gameContainer = document.getElementById('game');
const gameEndContainer = document.getElementById('game-end');

const cardsOnTable = document.querySelectorAll("[data-card]");
const cardImages = [];

for(var i = 1; i <= 52; i++){
    cardImages.push(i);
}

const resetDeal = function(){
    function shuffle(o) {
        for(var j, x, i = o.length; i; j = parseInt(Math.random() * i), x = o[--i], o[i] = o[j], o[j] = x);
        return o;
    };

    const randomCard = shuffle(cardImages);

    Array.prototype.forEach.call(cardsOnTable, function (card, index) {
        card.setAttribute("class","card-flip");
        window.setTimeout(setCards,800);
        function setCards() {
            card.setAttribute("style","background-image: url(images/"+[randomCard[index]]+".svg");            
        }
        console.log(card);
    });
};

declareWinner.hidden = true;
gameEndContainer.hidden = true;

document.getElementById('start-game').addEventListener('click', event => {
    startGame.hidden = true;
    declareWinner.hidden = false;

    say('Lets play poker');
    resetDeal();
    const numberOfPlayers = document.getElementById('player-count').value;

    if (window['WebSocket']) {
        const conn = new WebSocket('wss://' + document.location.host + '/ws');

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
            resetDeal();
        };

        conn.onopen = () => {
            conn.send(numberOfPlayers)
        }
    }
});

function say(msg) {
    window.speechSynthesis.speak(new SpeechSynthesisUtterance(msg));
}
