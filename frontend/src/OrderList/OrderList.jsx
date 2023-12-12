import {Order} from "../Order/Order";
import './styles.css'
export const OrderList = ({data}) => {

    return (
        <div className={"orderList-wrap"}>
        {
            data?.map((elem)=><Order key={elem.order_uid} data={elem}/>)
        }
        </div>
    )
}