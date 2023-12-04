function part01(input) {
        let cards = input.split("\n");
        let total = 0;
        for (const card of cards) {
                if (!card) continue;
                let [_, scratchPad] = card.split(": ");
                let [numbers, winningNumbers] = scratchPad.split(" | ");
                numbers = numbers.split(" ").filter((s) => s).map((n) => parseInt(n));
                winningNumbers = winningNumbers.split(" ").filter((s) => s).map((n) => parseInt(n));
                let currentPoints = 0;
                for (const number of numbers) {
                        if (!winningNumbers.includes(number)) continue;
                        if (currentPoints === 0) currentPoints = 1;
                        else currentPoints *= 2;
                }
                total += currentPoints;
        }
        return total;
}

module.exports = {
        part01
};
