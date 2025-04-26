import React, { createContext, useState, useEffect } from 'react';
import api from '../api';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            api.get('/me')
                .then((response) => {
                    setUser(response.data);
                    setLoading(false);
                })
                .catch(() => {
                    localStorage.removeItem('token');
                    setLoading(false);
                });
        } else {
            setLoading(false);
        }
    }, []);

    const login = async (username, password) => {
        try {
            const response = await api.post('/auth/login', { username, password });
            const { token } = response.data;
            localStorage.setItem('token', token);
            const userResponse = await api.get('/me');
            setUser(userResponse.data);
        } catch (error) {
            throw new Error(error.response?.data?.error || 'Ошибка входа');
        }
    };

    const register = async (username, password, balance) => {
        try {
            const response = await api.post('/auth/register', { username, password, balance });
            if (response.status === 201) {
                await login(username, password);
            }
        } catch (error) {
            throw new Error(error.response?.data?.error || 'Ошибка регистрации');
        }
    };

    const logout = () => {
        localStorage.removeItem('token');
        setUser(null);
    };

    const refreshUser = async () => {
        try {
            const response = await api.get('/me');
            setUser(response.data);
        } catch (error) {
            throw new Error(error.response?.data?.error || 'Ошибка обновления данных пользователя');
        }
    };

    return (
        <AuthContext.Provider value={{ user, login, register, logout, refreshUser, loading }}>
            {children}
        </AuthContext.Provider>
    );
};