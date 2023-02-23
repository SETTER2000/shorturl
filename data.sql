CREATE TABLE public.user
(
    id   VARCHAR(30) NOT NULL
);

CREATE TABLE public.shorturl
(
    slug    VARCHAR(30) NOT NULL,
    url     VARCHAR NOT NULL,
    user_id VARCHAR(30) NOT NULL
);

INSERT INTO public.user(id) VALUES ('1676935920173833222h_1');
INSERT INTO public.user(id) VALUES ('1676935920173833222h_2');
INSERT INTO public.user(id) VALUES ('1676935920173833222h_3');

INSERT INTO public.shorturl (slug, url, user_id) VALUES ('1676935920173833222h45','https://poaleell.com/chinese-crested/Poale-Ell-Adam','1676935920173833222h_1');
INSERT INTO public.shorturl (slug, url, user_id) VALUES ('1676935920173833222h46','https://poaleell.com/chinese-crested/Poale-Ell-Chen','1676935920173833222h_2');
INSERT INTO public.shorturl (slug, url, user_id) VALUES ('1676935920173833222h47','https://poaleell.com/chinese-crested/Poale-Ell-Cooper','1676935920173833222h_3');

SELECT * FROM public.user;
SELECT * FROM public.shorturl;