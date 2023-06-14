import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import authService from '../services/auth.service';

const Options = (props) => {
    return (
        <div className='mb-3'>
            <input type="text" className='form-control' id={"option" + props.id}></input>
        </div>
    );
}

const CreatePoll = () => {

    const navigate = useNavigate();

    const [isSubmitted, setIsSubmitted] = useState(false);
    const [optionsList, setOptionList] = useState([<Options id="1"/>, <Options id="2" />]);

    const handleAddOption = (event) => {
        setOptionList(prevlist => {
            return [...prevlist, <Options id={((prevlist.length) + 1)}/>]
        });
    }

    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitted(true);
        let options = []
        optionsList.map((x, i) => {
            let id = "option" + (i+1);
            console.log(event.target.elements);
            options.push(event.target.elements[id].value);
        })
        const requestBody = {
            "question": event.target.elements.question.value,
            "option": options,
        }
        try {
            const result = await authService.createPoll(requestBody);
            if (result.data) {
              navigate('/dashboard/allpolls');
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
                        <label htmlFor="question" className="form-label">Question</label>
                        <input type="text" className="form-control" id="question" />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="option1" className="form-label">Options</label>
                        {optionsList.map((x, i) => x)}
                    </div>
                    <button type="button" onClick={handleAddOption} className="btn btn-primary">Add option</button>
                    <button type="submit" disabled={isSubmitted} className="offset-1 btn btn-primary">Submit</button>
                </form>
            </div>
        </div>
    );
};

export default CreatePoll;