import React, { useState, useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import api from '../api';

const BetList = () => {
    const { user } = useContext(AuthContext);
    const [bets, setBets] = useState([]);
    const [error, setError] = useState('');

    const fetchBets = async () => {
        if (!user) {
            setError('Пожалуйста, войдите в систему');
            return;
        }

        try {
            setError('');
            const response = await api.get(`/bets/user/${user.id}`);
            console.log('Полученные ставки:', response.data);
            setBets(response.data);
        } catch (error) {
            setError('Ошибка при загрузке ставок: ' + (error.response?.data?.error || error.message));
        }
    };

    if (!user) {
        return <p>Пожалуйста, войдите в систему, чтобы просмотреть свои ставки.</p>;
    }

    return (
        <div>
            <h2>Ваши ставки (Пользователь: {user.username})</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <button onClick={fetchBets}>Показать мои ставки</button>
            <div>
                {bets.map((bet) => (
                    <div key={bet.id} className="bet">
                        ID ставки: {bet.id}
                        <br />
                        ID события: {bet.event_id}
                        <br />
                        Сумма: {bet.amount.toFixed(2)}
                        <br />
                        Исход: {bet.outcome}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default BetList;