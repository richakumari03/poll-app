import React from 'react';
import { Navigate, Outlet } from 'react-router-dom';
// import authService from '../services/auth.service';

const AuthGuard = () => {
    const authUser = true;//authService.getAuthUser();
    return authUser ? <Outlet /> : <Navigate to={'/'} replace />
}

export default AuthGuard;