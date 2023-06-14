import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import authService from '../services/auth.service';

export const Login = () => {

    const navigate = useNavigate();
    const [isSubmitted, setIsSubmitted] = useState(false);

    const handleSubmit = async (event) => {
        event.preventDefault();
        const requestBody = {
            "username": event.target.elements.username.value,
            "password": event.target.elements.password.value,
            "email": "",
        }
        setIsSubmitted(true)
        try {
          const result = await authService.login(requestBody);
          if (result.data) {
            console.log("entered");
            navigate('/dashboard');
          }
        } catch (error) {
          console.log(error);
        }
        setIsSubmitted(false)
    };

    return (
        <div className="row">
            <div className="col-6 offset-3">
                <form onSubmit={handleSubmit}>
                    <div className="mb-3">
                    <label htmlFor="username" className="form-label">Username</label>
                    <input type="text" className="form-control" id="username" />
                    </div>
                    <div className="mb-3">
                    <label htmlFor="password" className="form-label">Password</label>
                    <input type="password" className="form-control" id="password" />
                    </div>
                    <button type="submit" disabled={isSubmitted} className="btn btn-primary">Submit</button>
                </form>
            </div>
        </div>
    );
};

export const SignUp = () => {

    const navigate = useNavigate();
    const [isSubmitted, setIsSubmitted] = useState(false)

    const handleSubmit = async (event) => {
        event.preventDefault();
        const requestBody = {
            "username": event.target.elements.username.value,
            "email": event.target.elements.email.value,
            "password": event.target.elements.password.value
        }
        setIsSubmitted(true)
        try {
          const result = await authService.register(requestBody);
          if (result.data) {
            navigate('/');
          }
        } catch (error) {
            console.log(error);
        }
        setIsSubmitted(false)
    }

    return (
        <div className="row">
            <div className="col-6 offset-3">
                <form onSubmit={handleSubmit}>
                    <div className="mb-3">
                        <label htmlFor="username" className="form-label">Name</label>
                        <input type="text" className="form-control" id="username" required/>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="email" className="form-label">Email address</label>
                        <input type="email" className="form-control" id="email" required/>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="password" className="form-label">Password</label>
                        <input type="password" className="form-control" id="password" />
                    </div>
                    <button type="submit" className="btn btn-primary" disabled={isSubmitted}>Submit</button>
                </form>
            </div>
        </div>
    );
};
