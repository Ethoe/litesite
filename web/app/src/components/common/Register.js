import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { Container, Row, Col, Form, Button, Alert } from 'react-bootstrap';

function Register() {
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleFirstNameChange = (event) => {
        setFirstName(event.target.value);
    };

    const handleLastNameChange = (event) => {
        setLastName(event.target.value);
    };

    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        // Implement your registration logic here (similar to the login page)

        // For demonstration purposes, let's just show a success message
        setError('Registration successful!');

        navigate.push('/login'); // Replace with your desired URL
    };

    return (
        <Container>
            <Row className="justify-content-center mt-5">
                <Col xs={12} md={6}>
                    <h2>Register</h2>
                    <Form onSubmit={handleSubmit}>
                        <Form.Group as={Row}>
                            <Form.Label column sm="3">First Name:</Form.Label>
                            <Col sm="9">
                                <Form.Control type="text" value={firstName} onChange={handleFirstNameChange} />
                            </Col>
                        </Form.Group>
                        <Form.Group as={Row}>
                            <Form.Label column sm="3">Last Name:</Form.Label>
                            <Col sm="9">
                                <Form.Control type="text" value={lastName} onChange={handleLastNameChange} />
                            </Col>
                        </Form.Group>
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
                            Register
                        </Button>
                    </Form>
                    {error && <Alert variant="success" className="mt-3">{error}</Alert>}
                    <p className="mt-3">
                        Already have an account? <Link to="/login">Login</Link>
                    </p>
                </Col>
            </Row>
        </Container>
    );
}

export default Register;
