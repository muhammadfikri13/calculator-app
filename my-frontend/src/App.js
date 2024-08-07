import React, { useState} from 'react';
import './App.css';

function App() {
    const [operand1, setOperand1] = useState(0);
    const [operand2, setOperand2] = useState(0);
    const [operation, setOperation] = useState('add');
    const [result, setResult] = useState(null);

    const handleCalculate = async () => {
        const response = await fetch("http://localhost:9000/calculate", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ operation, operand1, operand2 }),
        });

        if (response.ok) {
            const data = await response.json();
            setResult(data.result);
        } else {
            alert('Error calculating result');
        }
    };

    return (
        <div className="App">
            <h1>Basic Calculator</h1>
            <input
                type="number"
                value={operand1}
                onChange={(e) => setOperand1(parseFloat(e.target.value))}
            />
            <select value={operation} onChange={(e) => setOperation(e.target.value)}>
                <option value="add">Tambah</option>
                <option value="subtract">Kurang</option>
                <option value="multiply">Kali</option>
                <option value="divide">Bagi</option>
            </select>
            <input
                type="number"
                value={operand2}
                onChange={(e) => setOperand2(parseFloat(e.target.value))}
            />
            <button onClick={handleCalculate}>Hitung</button>
            {result !== null &&  <h2>Hasil: {result}</h2>}
        </div>
    );
}

export default App;