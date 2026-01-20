const phoneticAlphabet = {
    'A': 'Alpha',
    'B': 'Bravo',
    'C': 'Charlie',
    'D': 'Delta',
    'E': 'Echo',
    'F': 'Foxtrot',
    'G': 'Golf',
    'H': 'Hotel',
    'I': 'India',
    'J': 'Juliett',
    'K': 'Kilo',
    'L': 'Lima',
    'M': 'Mike',
    'N': 'November',
    'O': 'Oscar',
    'P': 'Papa',
    'Q': 'Quebec',
    'R': 'Romeo',
    'S': 'Sierra',
    'T': 'Tango',
    'U': 'Uniform',
    'V': 'Victor',
    'W': 'Whiskey',
    'X': 'X-ray',
    'Y': 'Yankee',
    'Z': 'Zulu',
    '0': 'Zero',
    '1': 'One',
    '2': 'Two',
    '3': 'Three',
    '4': 'Four',
    '5': 'Five',
    '6': 'Six',
    '7': 'Seven',
    '8': 'Eight',
    '9': 'Nine'
};

const input = document.getElementById('nameInput');
const resultDiv = document.getElementById('result');

input.addEventListener('input', (e) => {
    const text = e.target.value.toUpperCase();
    resultDiv.innerHTML = '';

    for (let char of text) {
        if (char === ' ') {
             // Add a spacer for space
            const spacer = document.createElement('div');
            spacer.style.height = '10px';
            resultDiv.appendChild(spacer);
            continue;
        }

        const word = phoneticAlphabet[char];
        if (word) {
            const div = document.createElement('div');
            div.className = 'phonetic-word';
            div.innerHTML = `<span class="char">${char}</span> ${word}`;
            resultDiv.appendChild(div);
        } else {
             // Handle special characters if needed, or just display them
            const div = document.createElement('div');
            div.className = 'phonetic-word';
            div.innerHTML = `<span class="char">${char}</span> ${char}`;
            resultDiv.appendChild(div);
        }
    }
});
