import React from 'react';

import Dashboard from './pages/Dashboard';
import NotFound from './pages/NotFound';
import {Login, SignUp} from './pages/Register';
import CreatePoll from './pages/CreatePoll';
import { Routes, Route, Outlet } from 'react-router-dom';
import AuthGuard from './guards/auth.guard';
import AllPolls from './pages/AllPolls';

const AppRouter = () => {
    return (
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
          <Route element={<AuthGuard />}>
            <Route path="/dashboard" element={<Outlet />}>
              <Route index element={<Dashboard />} />
              <Route path="createPoll" element={<CreatePoll />}/>
              <Route path="allPolls" element={<AllPolls />}/>
            </Route>
            <Route path="*" element={<NotFound />} />
          </Route>
        </Routes>
    );
}

export default AppRouter;