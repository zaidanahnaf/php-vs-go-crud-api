<?php
header("Content-Type: application/json");
require "../conn.php";

$result = $conn->query("SELECT * FROM umkm");
$umkm = $result->fetch_all(MYSQLI_ASSOC);

echo json_encode($umkm);
?>
