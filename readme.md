category 的 update_time修改
update catogory set update_time=current_timestamp(0);



alter table user add column password varchar(20) not null;
update user set password=123456;

create table comment(
id bigint(20) primary key auto_increment not null,
content longtext not null,
article_id bigint(20) not null,
from_user bigint(20) not null,
at bigint(20) null,
create_time timestamp(0) not null default current_timestamp,
update_time timestamp(0) null default current_timestamp on update current_timestamp(0)
)

insert into comment (content, article_id, from_user) values('content1 from user1 article_id1', 1, 1)