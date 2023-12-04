function part02(input) {
        let cards = input.split("\n");
        let parsedCards = [];
        for (const card of cards) {
                if (!card) continue;
                let [details, scratchPad] = card.split(": ");
                let [_, cardNumberString] = details.replace(/ +/, " ").split(" ");
                let cardNumber = parseInt(cardNumberString);
                let [numbers, winningNumbers] = scratchPad.split(" | ");
                numbers = numbers.split(" ").filter((s) => s).map((n) => parseInt(n));
                winningNumbers = winningNumbers.split(" ").filter((s) => s).map((n) => parseInt(n));
                parsedCards[cardNumber - 1] = {
                        "total": 1,
                        "numbers": numbers,
                        "winningNumbers": winningNumbers,
                }
        }
        for (let i = 0; i < parsedCards.length; i++) {
                let matches = 0;
                for (const number of parsedCards[i]["numbers"]) {
                        if (parsedCards[i]["winningNumbers"].includes(number)) {
                                matches++;
                        }
                }
                for (let j = 1; j <= matches; j++) {
                        if (i+j < parsedCards.length) {
                                parsedCards[i+j]["total"] += parsedCards[i]["total"];
                        }
                }
        }
        return parsedCards.map((c) => c["total"]).reduce((p, c) => p+c);
}

module.exports = {
        part02
};
