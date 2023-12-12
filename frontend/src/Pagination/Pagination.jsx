import React, { useEffect, useState } from "react";

export const Pagination = ({ orderCounts, pageSize, setPage }) => {
    const [buttons, setButtons] = useState([1]);

    useEffect(() => {
        const totalPageCount = Math.ceil(orderCounts / pageSize);
        let myButtons = [];
        for (let i = 1; i < totalPageCount; i++) {
            myButtons.push(i);
        }
        setButtons(myButtons);
    }, [orderCounts]);

    const handleButtonClick = (pageNumber) => {
        console.log(pageNumber);
        setPage(pageNumber*10);
    };

    return (
        <div>
            {buttons.map((elem) => (
                <button key={elem} onClick={() => handleButtonClick(elem)}>
                    {elem}
                </button>
            ))}
        </div>
    );
};
