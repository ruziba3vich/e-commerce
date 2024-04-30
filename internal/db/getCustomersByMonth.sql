SELECT orders.order_id, customers.customer_country FROM Orders orders
INNER JOIN Customers customers ON customers.customer_id = orders.customer_id
INNER JOIN Order_details od ON od.order_id = orders.order_id
WHERE od.product_id = $1 AND EXTRACT(YEAR FROM orders.order_date) = $2
AND EXTRACT(MONTH FROM orders.order_date) = $3;
