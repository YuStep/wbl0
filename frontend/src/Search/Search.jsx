import React, { useState } from "react";

export const Search = ({ searchClick,toAllClick }) => {
    const [text, setText] = useState("");

    const handleChange = (e) => {
        setText(e.target.value);
    };

    const handleSearchClick = async () => {
        try {
            const response = await fetch(`/api/v1/order/${text}`);

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }

            let data = await response.json();

            const a = [data.data]

            console.log(a)
            // Обработка данных из ответа, например, вызов searchClick с данными
            searchClick({data:a})
        } catch (error) {
            console.error("Error fetching data:", error.message);
        }
    };

    const handleClick = () => {
        toAllClick()
    }
    return (
        <div className={"search-name"}>
            <input className={"search-input"} value={text} onChange={handleChange} />
            <button onClick={handleSearchClick}>Поиск</button>
            <button onClick={handleClick}>Список элементов</button>
        </div>
    );
};
