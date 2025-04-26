import React, { useState, useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import api from '../api';

const CreateBet = () => {
    const { user, refreshUser } = useContext(AuthContext);
    const [eventId, setEventId] = useState('');
    const [amount, setAmount] = useState('');
    const [outcome, setOutcome] = useState('win1');
    const [error, setError] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!user) {
            setError('Пожалуйста, войдите в систему');
            return;
        }

        try {
            setError('');
            const response = await api.post('/bets', {
                event_id: parseInt(eventId),
                amount: parseFloat(amount),
                outcome,
            });

            if (response.status === 201) {
                alert('Ставка успешно сделана!');
                await refreshUser(); // Обновляем данные пользователя (баланс)
                setEventId('');
                setAmount('');
                setOutcome('win1');
            }
        } catch (error) {
            setError('Ошибка при создании ставки: ' + (error.response?.data?.error || error.message));
        }
    };

    if (!user) {
        return <p>Пожалуйста, войдите в систему, чтобы сделать ставку.</p>;
    }

    return (
        <div>
            <h2>Сделать ставку (Пользователь: {user.username})</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <form onSubmit={handleSubmit}>
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