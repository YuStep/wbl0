package blanksDraft

import (
	"fmt"
	"github.com/YuStep/wbl0/internal/models"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateTable() error {

	_, err := r.db.Exec(`
        CREATE TABLE IF NOT EXISTS main (
            order_uid VARCHAR(255) PRIMARY KEY,
            track_number VARCHAR(255),
            entry VARCHAR(255),
            delivery_info JSONB,
            payment_info JSONB,
            items JSONB,
            locale VARCHAR(255),
            internal_signature VARCHAR(255),
  			customer_id VARCHAR(255),
  			delivery_service VARCHAR(255),
  			shardkey VARCHAR(255),
  			sm_id INTEGER,
            date_created VARCHAR(255),
            oof_shard VARCHAR(255)
        )
    `)

	return err
}

//func (r *OrderPostgres) SaveOrder(order models.Order) error {
//	deliveryJSON, err := json.Marshal(order.Delivery)
//	if err != nil {
//		return err
//	}
//	paymentJSON, err := json.Marshal(order.Payment)
//	if err != nil {
//		return err
//	}
//	itemsJSON, err := json.Marshal(order.Items)
//	if err != nil {
//		return err
//	}
//	_, err = r.db.Exec(`
//        INSERT INTO main (order_uid, track_number, entry, delivery_info, payment_info, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
//        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
//    `, order.OrderUID, order.TrackNumber, order.Entry, string(deliveryJSON), string(paymentJSON), string(itemsJSON), order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey, order.SmId, order.DateCreated, order.OofShard)
//
//	return err
//}

func (r *OrderPostgres) SaveOrder(order models.Order) error {
	_, err := r.db.Exec(`
        INSERT INTO main (order_uid, track_number, entry, delivery_info, payment_info, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
    `, order.OrderUID, order.TrackNumber, order.Entry, order.Delivery, order.Payment, order.Items, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey, order.SmId, order.DateCreated, order.OofShard)

	return err
}

func (r *OrderPostgres) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Select(&orders, `SELECT * FROM main;`)

	if err != nil {
		fmt.Printf("Error at getting all orders: %v\n", err)
		return nil, err
	}

	fmt.Println("Orders preloaded DB -> Cache")
	return orders, nil
}

//func (r *OrderPostgres) GetOrderById(id string) (*models.Order, error) {
//	var order *models.Order
//	query := fmt.Sprintf(`SELECT O.order_uid,O.track_number,O.entry,O.locale,O.internal_signature,
//    O.customer_id,O.delivery_service,O.shardkey,O.sm_id,O.date_created,O.oof_shard,D.name,D.phone,
//	D.zip,D.city,D.address,D.region,D.email,P.transaction,P.request_id,P.currency,P.provider,
//	P.amount,P.payment_dt,P.bank,P.delivery_cost,P.goods_total,P.custom_fee
//    FROM orders AS O JOIN delivery AS D ON D.order_id = O.order_uid JOIN payment AS P ON P.transaction = O.order_uid WHERE O.order_uid = $1`)
//
//	if err := r.db.Select(&order, query, id); err != nil {
//		return nil, err
//	}
//
//	var items []models.Item
//	query = fmt.Sprintf(`
//		SELECT items.chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status
//		FROM items JOIN public.orders_items AS oi on items.chrt_id = oi.chrt_id AND oi.order_id = $1
//		`)
//
//	if err := r.db.Select(&items, query, order.OrderUID); err != nil {
//		return nil, err
//	}
//
//	order.Items = items
//	return order, nil
//}
//
//func (r *OrderPostgres) GetOrders() ([]models.Order, error) {
//	var orders []models.Order
//	//	query := fmt.Sprintf(`SELECT
//	//    O.order_uid, O.track_number, O.entry, O.locale, O.internal_signature,
//	//    O.customer_id, O.delivery_service, O.shardkey, O.sm_id, O.date_created, O.oof_shard,
//	//    D.name , D.phone, D.zip , D.city,
//	//    D.address , D.region, D.email ,
//	//    P.transaction, P.request_id, P.currency, P.provider, P.amount, P.payment_dt, P.bank,
//	//    P.delivery_cost, P.goods_total, P.custom_fee,
//	//    I.chrt_id, I.track_number , I.price, I.rid, I.name ,
//	//    I.sale, I.size, I.total_price , I.nm_id, I.brand, I.status
//	//FROM
//	//    public.orders AS O
//	//JOIN
//	//    public.delivery AS D ON D.order_id = O.order_uid
//	//JOIN
//	//    public.payment AS P ON P.transaction = O.order_uid
//	//JOIN
//	//    public.orders_items AS OI ON OI.order_id = O.order_uid
//	//JOIN
//	//    public.items AS I ON I.chrt_id = OI.chrt_id`)
//	query := `SELECT json_agg("order") FROM
//    (SELECT order_uid, track_number, entry, locale, internal_signature,
//    customer_id,delivery_service, shardkey, sm_id, date_created, oof_shard,
//	(SELECT delivery FROM (SELECT D.name , D.phone, D.zip , D.city,
//    D.address , D.region, D.email FROM public.delivery AS D WHERE D.order_id = order_uid ) AS delivery),
//    (SELECT payment FROM (SELECT transaction, request_id, currency, provider, amount, payment_dt, public.banks.name AS bank,
//    delivery_cost, goods_total, custom_fee FROM public.payment JOIN banks ON public.payment.bank = public.banks.id
//    WHERE transaction = order_uid) AS payment),
//	(SELECT json_agg(row_to_json("items")) FROM (SELECT chrt_id,track_number,price,
//    rid,name,sale,size,total_price,nm_id,brand,status FROM public.items
//    WHERE public.items.order_id = order_uid) AS "items") AS "items"
//    FROM orders) AS "order"`
//
//	var buf []byte
//	if err := r.db.QueryRow(query).Scan(&buf); err != nil {
//		log.Fatal(err)
//	}
//	//var items []int
//	if err := json.Unmarshal(buf, &orders); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%#v\n", orders)
//
//	////Шаг 2 - добавление в Orders Items
//	//for i := range orders {
//	//	var items []models.Item
//	//	query := fmt.Sprintf(`
//	//	SELECT items.chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status
//	//	FROM items JOIN public.orders_items AS oi on items.chrt_id = oi.chrt_id AND oi.order_id = $1
//	//	`)
//	//
//	//	if err := r.db.Select(&items, query, orders[i].OrderUID); err != nil {
//	//		return nil, err
//	//	}
//	//	orders[i].Items = items
//	//}
//
//	return orders, nil
//}
