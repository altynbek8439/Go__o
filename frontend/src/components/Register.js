import React, { useState, useContext } from 'react';
import { AuthContext } from '../context/AuthContext';

const Register = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [balance, setBalance] = useState('');
    const [error, setError] = useState('');
    const { register } = useContext(AuthContext);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        try {
            await register(username, password, parseFloat(balance));
            setUsername('');
            setPassword('');
            setBalance('');
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <div>
            <h2>Регистрация</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Имя пользователя"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Пароль"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <input
                    type="number"
                    placeholder="Начальный баланс"
                    value={balance}
                    onChange={(e) => setBalance(e.target.value)}
                    step="0.1"
                    required
                />
                <button type="submit">Зарегистрироваться</button>
            </form>
        </div>
    );
};

export default Register;