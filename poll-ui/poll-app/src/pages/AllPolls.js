import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import authService from '../services/auth.service';

const PollQuestion = (props) => {

    const navigate = useNavigate();

    const [selectedOption, setSelectedOption] = useState(null);

    const handleOptionSelect = (optionId) => {
      setSelectedOption(optionId);
    };

    const handleVote = async() => {
        if (selectedOption !== null) {
            try {
                const requestBody = {
                    "questionId": props.Value.id,
                    "optionId": selectedOption
                }
                const result = await authService.updateVote(requestBody);
                if (result.data) {
                  window.location.reload();
                }
            } catch (error) {
                console.log(error);
            }
        }
    };

    return (
        <>
        <h3 className="mt-3">{props.QDesc}</h3>
        <div className="list-group">
          {props.Value.optionsDesc.map((option, index) => (
            <button
              key={index}
              type="button"
              className={`list-group-item list-group-item-action ${
                selectedOption === props.Value.optionsId[index] ? 'active' : ''
              }`}
              onClick={() => handleOptionSelect(props.Value.optionsId[index])}
            >
              {option}
            </button>
          ))}
        </div>
        <button
            type="button"
            className="btn btn-primary mt-3"
            onClick={handleVote}
            disabled={selectedOption === null}
        >
            Vote
        </button>
        </>
    );
}

const PollResults = (props) => {
    return (
        <>
            <h3 className="mt-5">{props.QDesc}</h3>
            <ul className="list-group">
                {props.Value.optionsDesc.map((option, index) => (
                <li
                    key={index}
                    className={`list-group-item d-flex justify-content-between align-items-center ${
                        props.Value.chosenOption === props.Value.optionsId[index] ? 'active' : ''
                      }`}
                >
                    {option}
                    <span className="badge bg-primary">{(props.Value.count[index] / props.Value.total) * 100} %</span>
                </li>
                ))}
            </ul>
        </>
    );
}

const AllPolls = () => {

    const [pollDetailsList, setpollDetailsList] = useState({});
    
    useEffect(() => {
        const fetchData = async () => {
          try {
            const result = await authService.getPolls();
            if(result.data) {
                const resp = {};
                const userMap = {};
                    const resultJsonArray = result.data.Result;
                    const userArray = result.data.UserResult;
                    if(userArray) {
                        userArray.forEach(element => {
                            userMap[element.QuestionId] = element.OptionId;
                        });
                    }
                    if(resultJsonArray) {
                        resultJsonArray.forEach(element => {
                            if(resp[element.QDesc] !== undefined) {
                                resp[element.QDesc].optionsId.push(element.OptionId);
                                resp[element.QDesc].optionsDesc.push(element.ODesc);
                                resp[element.QDesc].count.push(element.TotalCount);
                                resp[element.QDesc].total += element.TotalCount;
                            } else {
                                resp[element.QDesc] = {
                                    id: element.QuestionId,
                                    optionsId: [element.OptionId],
                                    optionsDesc: [element.ODesc],
                                    count: [element.TotalCount],
                                    total: element.TotalCount,
                                    chosenOption: userMap[element.QuestionId] ? userMap[element.QuestionId] : null
                                };
                            }
                            
                        });
                    }
                setpollDetailsList(resp);
            }
          } catch (error) {
            console.log('Error fetching data:', error);
          }
        };
    
        fetchData();
    }, []);

    return (
        <div className="container">
            {Object.entries(pollDetailsList).map(([key, value]) => {
                if(!value.chosenOption) {
                    return <PollQuestion QDesc={key} Value={value}/>
                }
            })}
            <h3 className="mt-4">Poll Results</h3>
            {Object.entries(pollDetailsList).map(([key, value]) => {
                if(value.chosenOption) {
                    return <PollResults QDesc={key} Value={value}/>
                }
            })}
        </div>
    );
};

export default AllPolls;