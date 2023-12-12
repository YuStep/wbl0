import React, { useEffect, useState } from "react";
import "./App.css";
import { OrderList } from "./OrderList/OrderList";
import { Pagination } from "./Pagination/Pagination";
import {Search} from "./Search/Search";

const AppHeader = ({setSearchData,setOneOrder}) => (
    <header>
      <Search searchClick={setSearchData} toAllClick={setOneOrder}/>
    </header>
);

const AppMain = ({ data, pageSize, setPage }) => {
    console.log(data)
    return (
        <main>
        {data && data.data.length > 0 ? (
            <>
                <OrderList data={data.data}/>
                <Pagination orderCounts={data.count} pageSize={pageSize} setPage={setPage}/>
            </>
        ) : (
            <p>No orders found.</p>
        )}
        </main>
    )
};

const AppFooter = () => (
    <footer>
      <h1>Footer</h1>
    </footer>
);

const App = () => {
  const [data, setData] = useState();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(undefined);
  const [page, setPage] = useState(1);
  const [limit, setLimit] = useState(10);
  const [oneOrder,setOneOrder] = useState(false)

  const setSearchData = (data) => {
      setData(data)
      if(oneOrder === true){
          return
      }
      setOneOrder(true)
  }

  const setOneOrderData = () => {
      if(oneOrder === false){
          return
      }
    setOneOrder(false)
  }

  useEffect(() => {
      if(oneOrder){
          return
      }
    const fetchData = async () => {
      setLoading(true);
      setError(null);

      try {
        const response = await fetch(`/api/v1/order?limit=${limit}&offset=${page}`);
        const responseJSON = await response.json();
        setData(prevData => ({
          ...prevData,
          data: responseJSON.data,
          count: responseJSON.count
        }));
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Unknown Error: get orders');
        setData(null);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [limit, page,oneOrder]);

  return (
      <div className="App">
        <AppHeader setSearchData={setSearchData} setOneOrder={setOneOrderData}/>
        {loading ? (
            <p>Loading...</p>
        ) : error ? (
            <p>Error: {error}</p>
        ) : (
            <AppMain data={data} pageSize={limit} setPage={setPage} />
        )}
        <AppFooter />
      </div>
  );
};

export default App;
