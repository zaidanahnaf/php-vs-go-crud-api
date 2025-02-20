<?php
header("Content-Type: application/json");
require "../conn.php";

$data = json_decode(file_get_contents("php://input"), true);
if (!isset($data["id"])) {
    echo json_encode(["error" => "ID diperlukan"]);
    exit;
}

$stmt = $conn->prepare("DELETE FROM umkm WHERE id=?");
$stmt->bind_param("i", $data["id"]);
$stmt->execute();

echo json_encode(["message" => "UMKM berhasil dihapus"]);
?>
