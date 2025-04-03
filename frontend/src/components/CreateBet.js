import React, { useState } from 'react';
import axios from 'axios';

const CreateBet = () => {
    const [userId, setUserId] = useState('');
    const [eventId, setEventId] = useState('');
    const [amount, setAmount] = useState('');
    const [outcome, setOutcome] = useState('win1');

    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log('Отправка ставки:', { user_id: parseInt(userId), event_id: parseInt(eventId), amount: parseFloat(amount), outcome });

        try {
            const response = await axios.post('http://localhost:8080/api/v1/bets', {
                user_id: parseInt(userId),
                event_id: parseInt(eventId),
                amount: parseFloat(amount),
                outcome,
            });

            if (response.status === 201) {
                alert('Ставка успешно сделана!');
                setUserId('');
                setEventId('');
                setAmount('');
                setOutcome('win1');
            }
        } catch (error) {
            console.error('Ошибка при создании ставки:', error);
        }
    };

    return (
        <div>
            <h2>Сделать ставку</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="number"
                    placeholder="ID пользователя"
                    value={userId}
                    onChange={(e) => setUserId(e.target.value)}
                    required
                />
                <input
                    type="number"
                    placeholder="ID события"
                    value={eventId}
                    onChange={(e) => setEventId(e.target.value)}
                    required
                />
                <input
                    type="number"
                    placeholder="Сумма ставки"
                    value={amount}
                    onChange={(e) => setAmount(e.target.value)}
                    step="0.1"
                    required
                />
                <select value={outcome} onChange={(e) => setOutcome(e.target.value)} required>
                    <option value="win1">Победа 1</option>
                    <option value="draw">Ничья</option>
                    <option value="win2">Победа 2</option>
                </select>
                <button type="submit">Сделать ставку</button>
            </form>
        </div>
    );
};

export default CreateBet;