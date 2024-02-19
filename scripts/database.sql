create database turismo_inventory;

create table tipos_atracciones(
	tip_atr_id smallint,
	tip_atr_nom varchar(50),
	tip_atr_fec_alt timestamp,
	tip_atr_fec_baj timestamp,
	primary key(tip_atr_id)
);
create sequence tip_atr_id_seq as smallint start 1 owned by tipos_atracciones.tip_atr_id;

create table atracciones (
	atr_id integer,
	atr_nom varchar(50),
	atr_dur interval,
	atr_cos decimal,
	atr_cap_max smallint,
	atr_tip_atr_id smallint,
	atr_fec_alt timestamp,
	atr_fec_act timestamp,
	atr_fec_baj timestamp,
	primary key(atr_id)
);
create sequence atr_id_seq as smallint start 1 owned by atracciones.atr_id
create index atr_idx_0 on atracciones(atr_tip_atr_id);

create table clientes(
	cli_id integer,
	cli_usrnm varchar(50),
	cli_din_disp decimal,
	cli_tiempo_disp interval,
	cli_tip_atr_id smallint,
	cli_fec_alt timestamp,
	cli_fec_act timestamp,
	cli_fec_baj timestamp,
	primary key(cli_id)
);
create sequence cli_id_seq as smallint start 1 owned by clientes.cli_id;
create unique index cli_usrnm_idx on clientes(cli_usrnm);

create table clientes_passwords(
	clpwd_cli_id integer,
	clpwd_usrnm varchar(50),
	clpwd_pwd varchar(255),
	clpwd_seed varchar(20),
	clpwd_fec_act timestamp,
	primary key(clpwd_cli_id)
);
alter table clientes_passwords
add constraint fk_clpwd_cli_id
foreign key (clpwd_cli_id) references clientes(cli_id);


create table promociones_porcentuales(
	ppor_id integer,
	ppor_nom varchar(50),
	ppor_desc decimal,
	ppor_fec_alt timestamp,
	ppor_fec_act timestamp,
	ppor_fec_baj timestamp,
	primary key(ppor_id)
);
create sequence ppor_id_seq as smallint start 1 owned by promociones_porcentuales.ppor_id;

create table rel_ppor_atr(
	rpporatr_ppor_id integer,
	rpporatr_atr_id integer,
	rpporatr_fec_alt timestamp,
	rpporatr_fec_baj timestamp,
	primary key(rpporatr_ppor_id, rpporatr_atr_id)
);
alter table rel_ppor_atr 
add constraint fk_rpporatr_ppor_id
foreign key (rpporatr_ppor_id) references promociones_porcentuales(ppor_id);
alter table rel_ppor_atr 
add constraint fk_rpporatr_atr_id
foreign key (rpporatr_atr_id) references atracciones(atr_id);

create table promociones_precio_fijo(
	ppf_id integer,
	ppf_nom varchar(50),
	ppf_mon decimal,
	ppf_fec_alt timestamp,
	ppf_fec_act timestamp,
	ppf_fec_baj timestamp,
	primary key (ppf_id)
);
create sequence ppf_id_seq as smallint start 1 owned by promociones_precio_fijo.ppf_id;

create table rel_ppf_atr(
	rppfatr_ppf_id integer,
	rppfatr_atr_id integer,
	rppfatr_fec_alt timestamp,
	rppfatr_fec_baj timestamp,
	primary key(rppfatr_ppf_id, rppfatr_atr_id)
);
alter table rel_ppf_atr 
add constraint fk_rppfatr_ppf_id
foreign key (rppfatr_ppf_id) references promociones_precio_fijo(ppf_id);
alter table rel_ppf_atr 
add constraint fk_rppfatr_atr_id
foreign key (rppfatr_atr_id) references atracciones(atr_id);