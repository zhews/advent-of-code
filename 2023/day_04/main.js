const fs = require("node:fs");

const { part01 } = require("./part_01");
const { part02 } = require("./part_02");

try {
        const input = fs.readFileSync("input.txt", "utf8");
        console.log(part01(input));
        console.log(part02(input));

} catch(error) {
        console.error(error);
}
