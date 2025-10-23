-- NO. 1
SELECT 
	COUNT(chdrnum) AS JUMLAH
FROM vm1dta_chdrpf
WHERE VALIDFLAG = 1;

-- NO. 2
SELECT
	vm1dta_descpf.longdesc AS "STATUS",
    COUNT(vm1dta_chdrpf.statcode) AS JUMLAH
FROM vm1dta_chdrpf
RIGHT JOIN vm1dta_descpf ON vm1dta_descpf.descitem = vm1dta_chdrpf.statcode
WHERE vm1dta_descpf.desctabl = "t3623"
GROUP BY vm1dta_descpf.longdesc;

-- NO. 3
SELECT 
	vm1dta_descpf.descitem AS KODE_PRODUK,
    vm1dta_descpf.longdesc AS NAMA_PRODUK,
    COUNT(vm1dta_chdrpf.cnttype) AS JUMLAH
FROM vm1dta_descpf
LEFT JOIN vm1dta_chdrpf ON vm1dta_chdrpf.cnttype = vm1dta_descpf.descitem
WHERE vm1dta_descpf.desctabl = 't5688' 
AND vm1dta_chdrpf.VALIDFLAG = 1
GROUP BY vm1dta_descpf.descitem, vm1dta_descpf.longdesc;

-- NO. 4
SELECT 
	DISTINCT vm1dta_agntpf.agntnum AS NOMOR_AGEN,
    vm1dta_clntpf.surname AS NAMA_AGEN,
    CONCAT_WS(
      ', ',
      NULLIF(vm1dta_clntpf.cltaddr01, ''),
      NULLIF(vm1dta_clntpf.cltaddr02, ''),
      NULLIF(vm1dta_clntpf.cltaddr03, ''),
      NULLIF(vm1dta_clntpf.cltaddr04, ''),
      NULLIF(vm1dta_clntpf.cltaddr05, '')
    ) AS ALAMAT_AGEN
FROM vm1dta_agntpf
INNER JOIN vm1dta_chdrpf ON vm1dta_chdrpf.agntnum = vm1dta_agntpf.agntnum
INNER JOIN vm1dta_clntpf ON vm1dta_clntpf.clntnum = vm1dta_agntpf.clntnum;

-- NO. 5
SELECT 
	polis.chdrnum AS NOMOR_POLIS,
    polis.cnttype AS KODE_PRODUK,
	produk1.longdesc AS NAMA_PRODUK,
    produk2.longdesc AS STATUS_PRODUK,
    polis.VALIDFLAG AS STATUS_VALID,
    client.surname AS NAMA_PEMEGANG_POLIS,
    CONCAT_WS(
      ', ',
      NULLIF(client.cltaddr01, ''),
      NULLIF(client.cltaddr02, ''),
      NULLIF(client.cltaddr03, ''),
      NULLIF(client.cltaddr04, ''),
      NULLIF(client.cltaddr05, '')
    ) AS  ALAMAT_PEMEGANG_POLIS,
    polis.agntnum AS NOMOR_AGEN,
    agen.surname AS NAMA_AGEN,
    CONCAT_WS(
      ', ',
      NULLIF(agen.cltaddr01, ''),
      NULLIF(agen.cltaddr02, ''),
      NULLIF(agen.cltaddr03, ''),
      NULLIF(agen.cltaddr04, ''),
      NULLIF(agen.cltaddr05, '')
    ) AS ALAMAT_AGEN
FROM vm1dta_chdrpf AS polis
INNER JOIN vm1dta_descpf AS produk1 ON produk1.descitem = polis.cnttype
INNER JOIN vm1dta_descpf AS produk2 ON produk2.descitem = polis.statcode
INNER JOIN vm1dta_agntpf ON vm1dta_agntpf.agntnum = polis.agntnum
INNER JOIN vm1dta_clntpf AS agen ON agen.clntnum = vm1dta_agntpf.clntnum
INNER JOIN vm1dta_clntpf AS client ON client.clntnum = polis.cownnum;
