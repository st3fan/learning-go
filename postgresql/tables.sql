
create table Things (
  ID serial not null unique,
  Name text not null
);

insert into Things (Name) values ('Toothbrush');
insert into Things (Name) values ('Radio');
insert into Things (Name) values ('Mirror');
insert into Things (Name) values ('Chair');
insert into Things (Name) values ('Table');
