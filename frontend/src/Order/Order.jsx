import './styles.css'

export const Order = ({data}) => {

    return (
        <div className={'order-wrap'}>
            <div className={'order-header'}>
                <h2>Заказ {data.order_uid}</h2>
                <p className={"order-p"}>
                    Трек номер {data.track_number}
                </p>
            </div>
            <div className={'order-info'}>
                <p className={"order-p"}>Адрес {data.delivery.address}</p>
                <p className={"order-p"}>Город {data.delivery.city}</p>
                <p className={"order-p"}>Email {data.delivery.email}</p>
                <p className={"order-p"}>Имя {data.delivery.name}</p>
                <p className={"order-p"}>Телефон {data.delivery.phone}</p>
            </div>
            <div>
                <select>
                    {data?.items.map((item)=>(
                        <option>
                            {item.brand} {item.name}
                        </option>))}
                </select>
            </div>
        </div>
    )
}