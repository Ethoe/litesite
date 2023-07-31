import React, { useState, useEffect } from 'react';
import apiClient from './../../services/apiClient';
import { useNavigate, Link } from 'react-router-dom';
import { Container, Row, Col, Form, Button, Alert } from 'react-bootstrap';

function Login({ setUser }) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    // Check if the user is already logged in on page load
    useEffect(() => {
        // Check the `session` cookie or use any other method to determine if the user is logged in
        // If the user is logged in, redirect to the dashboard or home page
        const isLoggedIn = document.cookie.includes('session=');
        if (isLoggedIn) {
            navigate.push('/'); // Replace with your desired URL
        }
    }, [navigate]);

    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        // Send login request to the backend
        apiClient
            .post('/login', { email, password })
            .then((response) => response.json())
            .then((data) => {
                // Handle login response
                if (data.success) {
                    // Redirect to dashboard or home page on successful login
                    // You can use React Router to handle the navigation
                    setUser(data.user)
                    navigate.pushState = '/'; // Replace with your desired URL
                } else {
                    setError('Invalid email or password');
                }
            })
            .catch((error) => {
                console.error('Error logging in:', error);
            });
    };

    return (
        <Container>
            <Row className="justify-content-center mt-5">
                <Col xs={12} md={6}>
                    <h2>Login</h2>
                    <Form onSubmit={handleSubmit}>
                        <Form.Group as={Row}>
                            <Form.Label column sm="3">Email:</Form.Label>
                            <Col sm="9">
                                <Form.Control type="email" value={email} onChange={handleEmailChange} />
                            </Col>
                        </Form.Group>
                        <Form.Group as={Row}>
                            <Form.Label column sm="3">Password:</Form.Label>
                            <Col sm="9">
                                <Form.Control type="password" value={password} onChange={handlePasswordChange} />
                            </Col>
                        </Form.Group>
                        <Button type="submit" style={{ backgroundColor: 'cornflowerblue', marginTop: '10px' }}>
                            Login
                        </Button>
                    </Form>
                    {error && <Alert variant="danger" className="mt-3">{error}</Alert>}
                    <p className="mt-3">
                        Don't have an account? <Link to="/register">Register</Link>
                    </p>
                </Col>
            </Row>
        </Container>
    );
}

export default Login;
