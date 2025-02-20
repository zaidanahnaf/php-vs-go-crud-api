<?php
header("Content-Type: application/json");
require "../conn.php";

$data = json_decode(file_get_contents("php://input"), true);
if (!isset($data["nama_usaha"]) || !isset($data["pemilik"]) || !isset($data["kategori"])) {
    echo json_encode(["error" => "Nama Usaha, Pemilik, dan Kategori diperlukan"]);
    exit;
}

$stmt = $conn->prepare("INSERT INTO umkm (nama_usaha, pemilik, kategori) VALUES (?, ?, ?)");
$stmt->bind_param("sss", $data["nama_usaha"], $data["pemilik"], $data["kategori"]);
$stmt->execute();

echo json_encode(["message" => "UMKM berhasil ditambahkan"]);
?>
