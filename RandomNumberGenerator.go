function getRandomNumber(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

const min = parseInt(prompt("Enter minimum value:"), 10);
const max = parseInt(prompt("Enter maximum value:"), 10);

console.log(`Random number between ${min} and ${max}: ${getRandomNumber(min, max)}`);
