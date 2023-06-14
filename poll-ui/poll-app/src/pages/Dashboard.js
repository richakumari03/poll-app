import React from 'react';
import { Link } from "react-router-dom";

const Dashboard = () => {

    return (
        <>
            <Link to="createPoll">
                <button className='btn btn-primary'>Create a Poll</button>
            </Link>
            <Link to="allPolls">
                <button className='offset-1 btn btn-primary'>My polls</button>
            </Link>
        </>
    );
};

export default Dashboard;