<?php
header("Content-Type: application/json");
require "../conn.php";

$data = json_decode(file_get_contents("php://input"), true);
if (!isset($data["id"]) || !isset($data["nama_usaha"]) || !isset($data["pemilik"]) || !isset($data["kategori"])) {
    echo json_encode(["error" => "ID, Nama Usaha, Pemilik, dan Kategori diperlukan"]);
    exit;
}

$stmt = $conn->prepare("UPDATE umkm SET nama_usaha=?, pemilik=?, kategori=? WHERE id=?");
$stmt->bind_param("sssi", $data["nama_usaha"], $data["pemilik"], $data["kategori"], $data["id"]);
$stmt->execute();

echo json_encode(["message" => "UMKM berhasil diperbarui"]);
?>
