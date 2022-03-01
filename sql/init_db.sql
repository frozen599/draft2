CREATE TABLE accounts (
    id serial PRIMARY KEY,
    pub_key text,
    priv_key text,
    amount real,
    address text,
    nonce int,
    CONSTRAINT pub_priv_key_unique UNIQUE (pub_key, priv_key)
);


INSERT INTO accounts (pub_key, priv_key, address, amount, nonce)
VALUES 
('04f6d4b5636f374a76fdacc9612d65e47a7bfecf3c17ac02a6318fa2815bb641b258ff79c1e5ccf3b13d394a764f1ca395ae879150eca67408000b3221f334a557', 
 '82e63b1ac40b27a84ce25ef17c865919586338447b70ebf0f037833d3872142d',
 '0x0b5253C570ae34E05F034A2836d0Cf09ebE10ECA',
 100,
 0),
('045c7c72b76f7f5097c61f405ca05539c38c4a0ce29086977544445a52a7840e7b42551d866dd3aa5af87278edce89b5ad270459a7d90763741555a83537ad6170',
 '0fb20ecbb26fee338c11eaf09440b06eb0c05086003e97a97e1f5a64e8cc9248',
 '0xFa502cEDEa20c926a6278D53592095aba87f7a4c',
 100,
 0),
('048d703b85a3dc37fc1a3e7696c4eb77d453e025780f19a94eac2650126096000a368b0dd45010c02b8f4499031a5e325fdced382a8d90b0f004f563f8dfbb43c1',
 'dcc96fd0a8395071c071b95e647157b774160cb91b269e3ada5eee41004ada31',
 '0x74b47c10D5883Db5F69839D26b6652352Bcd7E83',
 100,
 0);