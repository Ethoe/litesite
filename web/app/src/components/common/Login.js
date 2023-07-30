import React, { useState, useEffect } from 'react';
import apiClient from "./../../services/apiClient"

function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    // Check if the user is already logged in on page load
    useEffect(() => {
        // Check the `session` cookie or use any other method to determine if the user is logged in
        // If the user is logged in, redirect to the dashboard or home page
        const isLoggedIn = document.cookie.includes('session=');
        if (isLoggedIn) {
            window.location.href = '/'; // Replace with your desired URL
        }
    }, []);

    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        // Send login request to the backend
        apiClient.post('/login', { email, password })
            .then((response) => response.json())
            .then((data) => {
                // Handle login response
                if (data.success) {
                    // Redirect to dashboard or home page on successful login
                    // You can use React Router to handle the navigation
                    window.location.href = '/'; // Replace with your desired URL
                } else {
                    setError('Invalid email or password');
                }
            })
            .catch((error) => {
                console.error('Error logging in:', error);
            });
    };

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Email:</label>
                    <input type="email" value={email} onChange={handleEmailChange} />
                </div>
                <div>
                    <label>Password:</label>
                    <input type="password" value={password} onChange={handlePasswordChange} />
                </div>
                <button type="submit">Login</button>
                {error && <p>{error}</p>}
            </form>
        </div>
    );
}

export default Login;
