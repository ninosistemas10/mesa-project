-- Active: 1729800194520@@192.168.231.237@5432@deliveryDB
-- Active: 1697408358107@@192.168.1.135@5432@DBDelivery@public
-- Active: 1697408358107@@192.168.231.237@5432@DBDelivery@public
CREATE TABLE mesa(
	id				UUID			PRIMARY KEY,
	nombre			VARCHAR(50)		NOT NULL,
	url				VARCHAR(200)	NOT NULL,
	images			VARCHAR(250)	NOT NULL,
	activo			BOOLEAN			NOT NULL DEFAULT TRUE,
	created_at		INTEGER			NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at		INTEGER
);

CREATE TABLE users(
	id 		    	UUID 			DEFAULT uuid_generate_v4() PRIMARY KEY,
	nombre			VARCHAR(100)	NOT NULL,
	email			VARCHAR(254)	NOT NULL,
	password		VARCHAR(100)	NOT NULL,
	is_admin		BOOLEAN 		NOT NULL DEFAULT TRUE,
	images			JSONB			NOT NULL,
	details	    	JSONB 	    	NOT NULL,
	created_at		INTEGER			NOT NULL  DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at 		INTEGER,
	CONSTRAINT users_email_uk UNIQUE (email)
);

CREATE TABLE Promocion(
	id					UUID			PRIMARY KEY,
	nombre				VARCHAR(100)	NOT NULL,
	slug				VARCHAR(150)	GENERATE ALWAYS AS (lower(regexp_replace(nombre, '\s+', '-', 'g'))) STORED UNIQUE,
	description			TEXT			NOT NULL,
	image				VARCHAR(256)	NOT NULL,
	precio				NUMERIC(10,2)	NOT NULL CHECK (precio >= 0),
	features			JSONB			NOT NULL DEFAULT '{}'::JSONB,
	categoria			VARCHAR(100)	NOT NULL,
	stock_disponible	INTEGER			NOT NULL CHECK(stock_disponible >= 0) DEFAULT 0,
	activo				BOOLEAN			NOT NULL,
	created_at			INTEGER			NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at			INTEGER			
);

CREATE TABLE Promocion_Productos(
	id				UUID			PRIMARY KEY,
	promocionId		UUID			NOT NULL Promocion(id) ON DELETE CASCADE,
	productoId		UUID			NOT NULL Producto(id) ON DELETE RESTRICT,
	cantidad		INTEGER			NOT NULL CHECK(cantidad > 0),
	created_at		INTEGER			NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	update_at		INTEGER	,		
    FOREIGN KEY(promocionId) REFERENCES Promocion(id),
    FOREIGN KEY (productoId) REFERENCES Productos(id)
);


CREATE TABLE Category(
    id              UUID            PRIMARY KEY,
    nombre		    VARCHAR(100)    NOT NULL UNIQUE,
    description     TEXT            NOT NULL,
	images			VARCHAR(250)	NOT NULL,
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at      INTEGER
);

CREATE TABLE Productos (
    id            UUID            PRIMARY KEY,
    idcategoria  UUID            NOT NULL,
    nombre        VARCHAR(128)    NOT NULL UNIQUE,
    precio        NUMERIC(10,2)   NOT NULL,
    imagen        VARCHAR(250)    NOT NULL,
    descripcion   TEXT            NOT NULL,
    activo        BOOLEAN         NOT NULL DEFAULT TRUE,
    time          TIMESTAMP       NOT NULL,
    calorias      NUMERIC(6,2)    NOT NULL,
    created_at    INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at    INTEGER,
    FOREIGN KEY(idcategoria) REFERENCES Category(id)
);
