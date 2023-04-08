# Cudo Skill Test
Please read info below to run every project

## get-matrix-element-by-coordinate
Just execute `go run main.go` and follow the input text command

## order-report-chart
This is a back-end service run on the TCP port 8000, execute `go run .` to run the service. Then on your browser hit `localhost:8000`, you should see the result

### Another Info
SQL to Create Order Report Data.<br>
*Note: This project is using tables from W3Schools*
```sql
SELECT OrderDate, CategoryName, SUM(Quantity) AS Quantity FROM OrderDetails JOIN Orders ON OrderDetails.OrderID = Orders.OrderID JOIN Products ON OrderDetails.ProductID = Products.ProductID JOIN Categories ON Products.CategoryID = Categories.CategoryID GROUP BY OrderDate, CategoryName ORDER BY OrderDate ASC;
```
