import React, { useState } from 'react';
import axios from 'axios';

const CreateEvent = ({ onEventCreated }) => {
    const [name, setName] = useState('');
    const [date, setDate] = useState('');
    const [oddsWin1, setOddsWin1] = useState('');
    const [oddsDraw, setOddsDraw] = useState('');
    const [oddsWin2, setOddsWin2] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const response = await axios.post('http://localhost:8080/api/v1/events', {
                name,
                date,
                odds_win1: parseFloat(oddsWin1),
                odds_draw: parseFloat(oddsDraw),
                odds_win2: parseFloat(oddsWin2),
            });

            if (response.status === 201) {
                onEventCreated();
                setName('');
                setDate('');
                setOddsWin1('');
                setOddsDraw('');
                setOddsWin2('');
            }
        } catch (error) {
            console.error('Ошибка при создании события:', error);
        }
    };

    return (
        <div>
            <h2>Добавить событие</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Название события"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                />
                <input
                    type="date"
                    value={date}
                    onChange={(e) => setDate(e.target.value)}
                    required
                />
                <input
                    type="number"
                    placeholder="Коэффициент на победу 1"
                    value={oddsWin1}
                    onChange={(e) => setOddsWin1(e.target.value)}
                    step="0.1"
                    required
                />
                <input
                    type="number"
                    placeholder="Коэффициент на ничью"
                    value={oddsDraw}
                    onChange={(e) => setOddsDraw(e.target.value)}
                    step="0.1"
                    required
                />
                <input
                    type="number"
                    placeholder="Коэффициент на победу 2"
                    value={oddsWin2}
                    onChange={(e) => setOddsWin2(e.target.value)}
                    step="0.1"
                    required
                />
                <button type="submit">Добавить событие</button>
            </form>
        </div>
    );
};

export default CreateEvent;