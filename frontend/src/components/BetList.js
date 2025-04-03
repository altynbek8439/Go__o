import React, { useState } from 'react';
import axios from 'axios';

const BetList = () => {
    const [userId, setUserId] = useState('');
    const [bets, setBets] = useState([]);

    const fetchBets = async () => {
        if (!userId) {
            alert('Введите ID пользователя');
            return;
        }

        try {
            const response = await axios.get(`http://localhost:8080/api/v1/bets/user/${userId}`);
            console.log('Полученные ставки:', response.data);
            setBets(response.data);
        } catch (error) {
            console.error('Ошибка при загрузке ставок:', error);
        }
    };

    return (
        <div>
            <h2>Ставки пользователя</h2>
            <input
                type="number"
                placeholder="ID пользователя"
                value={userId}
                onChange={(e) => setUserId(e.target.value)}
            />
            <button onClick={fetchBets}>Показать ставки</button>
            <div>
                {bets.map(bet => (
                    <div key={bet.id} className="bet">
                        ID ставки: {bet.id}<br />
                        ID события: {bet.event_id}<br />
                        Сумма: {bet.amount.toFixed(2)}<br />
                        Исход: {bet.outcome}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default BetList;