load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:TEgegKbBqByGUb1Coo1pc2qIdf2xw6v0mYyLSYtyopE=",
        version = "v0.0.1-2019.2.2",
    )
    go_repository(
        name = "com_github_agnivade_levenshtein",
        importpath = "github.com/agnivade/levenshtein",
        sum = "h1:3oJU7J3FGFmyhn8KHjmVaZCN5hxTr7GxgRue+sxIXdQ=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_ajg_form",
        importpath = "github.com/ajg/form",
        sum = "h1:zvClvFQwU++UpIUBGC8YmDlfhUrweEy1R1Fj1gu5iIM=",
        version = "v0.0.0-20160822230020-523a5da1a92f",
    )
    go_repository(
        name = "com_github_alecthomas_template",
        importpath = "github.com/alecthomas/template",
        sum = "h1:cAKDfWh5VpdgMhJosfJnn5/FoN2SRZ4p7fJNX58YPaU=",
        version = "v0.0.0-20160405071501-a0175ee3bccc",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        importpath = "github.com/alecthomas/units",
        sum = "h1:qet1QNfXsQxTZqLG4oE62mJzwPIB8+Tee4RNCL9ulrY=",
        version = "v0.0.0-20151022065526-2efee857e7cf",
    )
    go_repository(
        name = "com_github_andreyvit_diff",
        importpath = "github.com/andreyvit/diff",
        sum = "h1:bvNMNQO63//z+xNgfBlViaCIJKLlCJ6/fmUseuG0wVQ=",
        version = "v0.0.0-20170406064948-c7f18ee00883",
    )
    go_repository(
        name = "com_github_andybalholm_cascadia",
        importpath = "github.com/andybalholm/cascadia",
        sum = "h1:BuuO6sSfQNFRu1LppgbD25Hr2vLYW25JvxHs5zzsLTo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_apache_thrift",
        importpath = "github.com/apache/thrift",
        sum = "h1:pODnxUFNcjP9UTLZGTdeh+j16A8lJbRvD3rOtrk/7bs=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_armon_consul_api",
        importpath = "github.com/armon/consul-api",
        sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
        version = "v0.0.0-20180202201655-eb2c6b5be1b6",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        importpath = "github.com/asaskevich/govalidator",
        sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
        version = "v0.0.0-20190424111038-f61b66f89f4a",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:/4+rDPe0W95KBmNGYCG+NUvdL8ssPYBMxL+aSCg6nIA=",
        version = "v1.17.7",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
        version = "v0.0.0-20170929234023-d6e3b3328b78",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bitly_go_hostpool",
        importpath = "github.com/bitly/go-hostpool",
        sum = "h1:mXoPYz/Ul5HYEDvkta6I8/rnYM5gSdSV2tJ6XbZuEtY=",
        version = "v0.0.0-20171023180738-a3a6125de932",
    )
    go_repository(
        name = "com_github_bmizerany_assert",
        importpath = "github.com/bmizerany/assert",
        sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
        version = "v0.0.0-20160611221934-b7ed37b82869",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_burntsushi_xgb",
        importpath = "github.com/BurntSushi/xgb",
        sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version = "v0.0.0-20160522181843-27f122750802",
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        importpath = "github.com/cenkalti/backoff",
        sum = "h1:tKJnvO2kl0zmb/jA5UKAt4VoEVw1qxKWjE/Bpp46npY=",
        version = "v2.1.1+incompatible",
    )
    go_repository(
        name = "com_github_cespare_xxhash",
        importpath = "github.com/cespare/xxhash",
        sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cockroachdb_apd",
        importpath = "github.com/cockroachdb/apd",
        sum = "h1:3LFP3629v+1aKXU5Q37mxmRxX/pIu1nijXydLShEq5I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cockroachdb_cockroach_go",
        importpath = "github.com/cockroachdb/cockroach-go",
        sum = "h1:2zRrJWIt/f9c9HhNHAgrRgq0San5gRRUJTBXLkchal0=",
        version = "v0.0.0-20181001143604-e0a95dfd547c",
    )
    go_repository(
        name = "com_github_codahale_hdrhistogram",
        importpath = "github.com/codahale/hdrhistogram",
        sum = "h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=",
        version = "v0.0.0-20161010025455-3a0bb77429bd",
    )
    go_repository(
        name = "com_github_codegangsta_negroni",
        importpath = "github.com/codegangsta/negroni",
        sum = "h1:+aYywywx4bnKXWvoWtRfJ91vC59NbEhEY03sZjQhbVY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_containerd",
        importpath = "github.com/containerd/containerd",
        sum = "h1:8lqLbl7u1j3MmiL9cJ/O275crSq7bfwUayvvatEupQk=",
        version = "v1.2.7",
    )
    go_repository(
        name = "com_github_containerd_continuity",
        importpath = "github.com/containerd/continuity",
        sum = "h1:4BX8f882bXEDKfWIf0wa8HRvpnBoPszJJXL+TVbBw4M=",
        version = "v0.0.0-20181203112020-004b46473808",
    )
    go_repository(
        name = "com_github_corbym_gocrest",
        importpath = "github.com/corbym/gocrest",
        sum = "h1:gwEdq6RkTmq+09CTuM29DfKOCtZ7G7bcyxs3IZ6EVdU=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_coreos_bbolt",
        importpath = "github.com/coreos/bbolt",
        sum = "h1:n6AiVyVRKQFNb6mJlwESEvvLoDyiTzXX7ORAUlkeBdY=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_github_coreos_etcd",
        importpath = "github.com/coreos/etcd",
        sum = "h1:8F3hqu9fGYLBifCmRCJsicFqDx/D68Rt3q1JMazcgBQ=",
        version = "v3.3.13+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_etcd",
        importpath = "github.com/coreos/go-etcd",
        sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_oidc",
        importpath = "github.com/coreos/go-oidc",
        sum = "h1:+RStIopZ8wooMx+Vs5Bt8zMXxV1ABl5LbakNExNmZIg=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        importpath = "github.com/coreos/go-systemd",
        sum = "h1:F5haX7vjVVG0kc13fIWeqUViNPyEJxv/OmvnBo0Yme4=",
        version = "v0.0.0-20190321100706-95778dfbb74e",
    )
    go_repository(
        name = "com_github_coreos_pkg",
        importpath = "github.com/coreos/pkg",
        sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
        version = "v0.0.0-20180928190104-399ea9e2e55f",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man",
        importpath = "github.com/cpuguy83/go-md2man",
        sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
        version = "v1.0.10",
    )
    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:6pwm8kMQKCmgUg0ZHTm5+/YvRK0s3THD/28+T6/kk4A=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_cznic_b",
        importpath = "github.com/cznic/b",
        sum = "h1:UHFGPvSxX4C4YBApSPvmUfL8tTvWLj2ryqvT9K4Jcuk=",
        version = "v0.0.0-20180115125044-35e9bbe41f07",
    )
    go_repository(
        name = "com_github_cznic_fileutil",
        importpath = "github.com/cznic/fileutil",
        sum = "h1:7uSNgsgcarNk4oiN/nNkO0J7KAjlsF5Yv5Gf/tFdHas=",
        version = "v0.0.0-20180108211300-6a051e75936f",
    )
    go_repository(
        name = "com_github_cznic_golex",
        importpath = "github.com/cznic/golex",
        sum = "h1:CVAqftqbj+exlab+8KJQrE+kNIVlQfJt58j4GxCMF1s=",
        version = "v0.0.0-20170803123110-4ab7c5e190e4",
    )
    go_repository(
        name = "com_github_cznic_internal",
        importpath = "github.com/cznic/internal",
        sum = "h1:FHpbUtp2K8X53/b4aFNj4my5n+i3x+CQCZWNuHWH/+E=",
        version = "v0.0.0-20180608152220-f44710a21d00",
    )
    go_repository(
        name = "com_github_cznic_lldb",
        importpath = "github.com/cznic/lldb",
        sum = "h1:AIA+ham6TSJ+XkMe8imQ/g8KPzMUVWAwqUQQdtuMsHs=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cznic_mathutil",
        importpath = "github.com/cznic/mathutil",
        sum = "h1:XNT/Zf5l++1Pyg08/HV04ppB0gKxAqtZQBRYiYrUuYk=",
        version = "v0.0.0-20180504122225-ca4c9f2c1369",
    )
    go_repository(
        name = "com_github_cznic_ql",
        importpath = "github.com/cznic/ql",
        sum = "h1:lcKp95ZtdF0XkWhGnVIXGF8dVD2X+ClS08tglKtf+ak=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_cznic_sortutil",
        importpath = "github.com/cznic/sortutil",
        sum = "h1:hxuZop6tSoOi0sxFzoGGYdRqNrPubyaIf9KoBG9tPiE=",
        version = "v0.0.0-20150617083342-4c7342852e65",
    )
    go_repository(
        name = "com_github_cznic_strutil",
        importpath = "github.com/cznic/strutil",
        sum = "h1:0rkFMAbn5KBKNpJyHQ6Prb95vIKanmAe62KxsrN+sqA=",
        version = "v0.0.0-20171016134553-529a34b1c186",
    )
    go_repository(
        name = "com_github_cznic_zappy",
        importpath = "github.com/cznic/zappy",
        sum = "h1:YKKpTb2BrXN2GYyGaygIdis1vXbE7SSAG9axGWIMClg=",
        version = "v0.0.0-20160723133515-2533cb5b45cc",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_denisenkom_go_mssqldb",
        importpath = "github.com/denisenkom/go-mssqldb",
        sum = "h1:tkum0XDgfR0jcVVXuTsYv/erY2NnEDqwRojbxR1rBYA=",
        version = "v0.0.0-20190515213511-eb9f6a1743f3",
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        importpath = "github.com/dgrijalva/jwt-go",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_dgryski_go_sip13",
        importpath = "github.com/dgryski/go-sip13",
        sum = "h1:Yqiad0+sloMPdd/0Fg22actpFx0dekpzt1xJmVNVkU0=",
        version = "v0.0.0-20190329191031-25c5027a8c7b",
    )
    go_repository(
        name = "com_github_dhui_dktest",
        importpath = "github.com/dhui/dktest",
        sum = "h1:kwX5a7EkLcjo7VpsPQSYJcKGbXBXdjI9FGjuUj1jn6I=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_docker_distribution",
        importpath = "github.com/docker/distribution",
        sum = "h1:neUDAlf3wX6Ml4HdqTrbcOHXtfRN0TFIwt6YFL7N9RU=",
        version = "v2.7.0+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        sum = "h1:mzrx39dGtGq0VEnTHjnakmczd4uFbhx2cZU3BJDsLdc=",
        version = "v0.7.3-0.20190817195342-4760db040282",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_eapache_go_resiliency",
        importpath = "github.com/eapache/go-resiliency",
        sum = "h1:1NtRmCAqadE2FN4ZcN6g90TP3uk8cg9rn9eNK2197aU=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_eapache_go_xerial_snappy",
        importpath = "github.com/eapache/go-xerial-snappy",
        sum = "h1:YEetp8/yCZMuEPMUDHG0CW/brkkEp8mzqk2+ODEitlw=",
        version = "v0.0.0-20180814174437-776d5712da21",
    )
    go_repository(
        name = "com_github_eapache_queue",
        importpath = "github.com/eapache/queue",
        sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:aaQcKT9WumO6JEJcRyTqFVq4XUZiUcKR2/GI31TOcz8=",
        version = "v0.0.0-20170320065105-0bce6a688712",
    )
    go_repository(
        name = "com_github_elazarl_goproxy",
        importpath = "github.com/elazarl/goproxy",
        sum = "h1:LCoguo7Zd0MByKMbQbTvcZw7HiBcbvew+MOcwsJVwrY=",
        version = "v0.0.0-20181003060214-f58a169a71a5",
    )
    go_repository(
        name = "com_github_erikstmartin_go_testdb",
        importpath = "github.com/erikstmartin/go-testdb",
        sum = "h1:Yzb9+7DPaBjB8zlTR87/ElzFsnQfuHnVUVqpZZIcV5Y=",
        version = "v0.0.0-20160219214506-8d10e4a1bae5",
    )
    go_repository(
        name = "com_github_fasthttp_contrib_websocket",
        importpath = "github.com/fasthttp-contrib/websocket",
        sum = "h1:DddqAaWDpywytcG8w/qoQ5sAN8X12d3Z3koB0C3Rxsc=",
        version = "v0.0.0-20160511215533-1f3b11f56072",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_fatih_structs",
        importpath = "github.com/fatih/structs",
        sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
        version = "v1.4.7",
    )
    go_repository(
        name = "com_github_fsouza_fake_gcs_server",
        importpath = "github.com/fsouza/fake-gcs-server",
        sum = "h1:Un0BXUXrRWYSmYyC1Rqm2e2WJfTPyDy/HGMz31emTi8=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_gavv_httpexpect",
        importpath = "github.com/gavv/httpexpect",
        sum = "h1:1X9kcRshkSKEjNJJxX9Y9mQ5BRfbxU5kORdjhlA1yX8=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_globalsign_mgo",
        importpath = "github.com/globalsign/mgo",
        sum = "h1:DujepqpGd1hyOd7aW59XpK7Qymp8iy83xq74fLr21is=",
        version = "v0.0.0-20181015135952-eeefdecb41b8",
    )
    go_repository(
        name = "com_github_go_bindata_go_bindata",
        importpath = "github.com/go-bindata/go-bindata",
        sum = "h1:tR4f0e4VTO7LK6B2YWyAoVEzG9ByG1wrXB4TL9+jiYg=",
        version = "v3.1.1+incompatible",
    )
    go_repository(
        name = "com_github_go_chi_chi",
        importpath = "github.com/go-chi/chi",
        sum = "h1:maB6vn6FqCxrpz4FqWdh4+lwpyZIQS7YEAUcHlgXVRs=",
        version = "v4.0.2+incompatible",
    )
    go_repository(
        name = "com_github_go_chi_cors",
        importpath = "github.com/go-chi/cors",
        sum = "h1:e6x8k7uWbUwYs+aXDoiUzeQFT6l0cygBYyNhD7/1Tg0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_chi_jwtauth",
        importpath = "github.com/go-chi/jwtauth",
        sum = "h1:hPhobLUgh7fMpA1qUDdId14u2Z93M22fCNPMVLNWeHU=",
        version = "v4.0.3+incompatible",
    )
    go_repository(
        name = "com_github_go_chi_render",
        importpath = "github.com/go-chi/render",
        sum = "h1:4/5tis2cKaNdnv9zFLfXzcquC9HbeZgCnxGnKrltBS8=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_go_kit_kit",
        importpath = "github.com/go-kit/kit",
        sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:MP4Eh7ZCb31lleYCFuwm0oe4/YGak+5l1vA2NOE80nA=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_go_mail_mail",
        importpath = "github.com/go-mail/mail",
        sum = "h1:UzNOn0k5lpfVtO31cK3hn6I4VEVGhe3lX8AJBAxXExM=",
        version = "v2.3.1+incompatible",
    )
    go_repository(
        name = "com_github_go_openapi_analysis",
        importpath = "github.com/go-openapi/analysis",
        sum = "h1:5Z7zgx/EAmE9bf7cuTU1hkGQlZGzGDf9a1m57xRfNZk=",
        version = "v0.19.6",
    )
    go_repository(
        name = "com_github_go_openapi_errors",
        importpath = "github.com/go-openapi/errors",
        sum = "h1:a2kIyV3w+OS3S97zxUndRVD46+FhGOUBDFY7nmu4CsY=",
        version = "v0.19.2",
    )
    go_repository(
        name = "com_github_go_openapi_inflect",
        importpath = "github.com/go-openapi/inflect",
        sum = "h1:9jCH9scKIbHeV9m12SmPilScz6krDxKRasNNSNPXu/4=",
        version = "v0.19.0",
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        importpath = "github.com/go-openapi/jsonpointer",
        sum = "h1:gihV7YNZK1iK6Tgwwsxo2rJbD1GTbdm72325Bq8FI3w=",
        version = "v0.19.3",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        importpath = "github.com/go-openapi/jsonreference",
        sum = "h1:5cxNfTy0UVC3X8JL5ymxzyoUZmo8iZb+jeTWn7tUa8o=",
        version = "v0.19.3",
    )
    go_repository(
        name = "com_github_go_openapi_loads",
        importpath = "github.com/go-openapi/loads",
        sum = "h1:5I4CCSqoWzT+82bBkNIvmLc0UOsoKKQ4Fz+3VxOB7SY=",
        version = "v0.19.4",
    )
    go_repository(
        name = "com_github_go_openapi_runtime",
        importpath = "github.com/go-openapi/runtime",
        sum = "h1:la3rOpJ+UnJSCgsgMv/xKHCMCpKJtFVa7ai4mPnV2ts=",
        version = "v0.19.8",
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        importpath = "github.com/go-openapi/spec",
        sum = "h1:ixzUSnHTd6hCemgtAJgluaTSGYpLNpJY4mA2DIkdOAo=",
        version = "v0.19.4",
    )
    go_repository(
        name = "com_github_go_openapi_strfmt",
        importpath = "github.com/go-openapi/strfmt",
        sum = "h1:eRfyY5SkaNJCAwmmMcADjY31ow9+N7MCLW7oRkbsINA=",
        version = "v0.19.3",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        importpath = "github.com/go-openapi/swag",
        sum = "h1:lTz6Ys4CmqqCQmZPBlbQENR1/GucA2bzYTE12Pw4tFY=",
        version = "v0.19.5",
    )
    go_repository(
        name = "com_github_go_openapi_validate",
        importpath = "github.com/go-openapi/validate",
        sum = "h1:QhCBKRYqZR+SKo4gl1lPhPahope8/RLt6EVgY8X80w0=",
        version = "v0.19.5",
    )
    go_repository(
        name = "com_github_go_ozzo_ozzo_validation",
        importpath = "github.com/go-ozzo/ozzo-validation",
        sum = "h1:msy24VGS42fKO9K1vLz82/GeYW1cILu7Nuuj1N3BBkE=",
        version = "v3.6.0+incompatible",
    )
    go_repository(
        name = "com_github_go_playground_locales",
        importpath = "github.com/go-playground/locales",
        sum = "h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_go_playground_universal_translator",
        importpath = "github.com/go-playground/universal-translator",
        sum = "h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=",
        version = "v0.17.0",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:g24URVg0OFbNUTx9qqY1IRZ9D9z3iPyi5zKhQZpNwpA=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_go_swagger_go_swagger",
        importpath = "github.com/go-swagger/go-swagger",
        sum = "h1:4ejFC3gnhLTI6S/j+S+7GfEaVBV3BBI2IEcFVHUteIQ=",
        version = "v0.20.0",
    )
    go_repository(
        name = "com_github_go_swagger_scan_repo_boundary",
        importpath = "github.com/go-swagger/scan-repo-boundary",
        sum = "h1:l9rI6sNaZgNC0LnF3MiE+qTmyBA/tZAg1rtyrGbUMK0=",
        version = "v0.0.0-20180623220736-973b3573c013",
    )
    go_repository(
        name = "com_github_gobuffalo_buffalo",
        importpath = "github.com/gobuffalo/buffalo",
        sum = "h1:Fyn55HJULJpFPMUNx9lrPK31qvr37+bpNGFbpAOauGI=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_gobuffalo_buffalo_plugins",
        importpath = "github.com/gobuffalo/buffalo-plugins",
        sum = "h1:yZ6USaSdAKpogRS8DZJgeG7/CTPGmyhplwifphmmegw=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_gobuffalo_buffalo_pop",
        importpath = "github.com/gobuffalo/buffalo-pop",
        sum = "h1:8aXdBlo9MEKFGHyl489+28Jw7Ud59Th1U+5Ayu1wNL0=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_gobuffalo_envy",
        importpath = "github.com/gobuffalo/envy",
        sum = "h1:OsV5vOpHYUpP7ZLS6sem1y40/lNX1BZj+ynMiRi21lQ=",
        version = "v1.6.15",
    )
    go_repository(
        name = "com_github_gobuffalo_events",
        importpath = "github.com/gobuffalo/events",
        sum = "h1:ukq5ys/h0TuiX7eLJyZBD1dJOy0r19JTEYmgXKG9j+Y=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_gobuffalo_fizz",
        importpath = "github.com/gobuffalo/fizz",
        sum = "h1:JJOkmlStog5AiBL434UoGMJ896p3MnTnzedFVaZSF3k=",
        version = "v1.0.12",
    )
    go_repository(
        name = "com_github_gobuffalo_flect",
        importpath = "github.com/gobuffalo/flect",
        sum = "h1:HZOs07hF3AmoaUj4HJQHV5RqfOuGnPZI7aFcireIrww=",
        version = "v0.0.0-20190117212819-a62e61d96794",
    )
    go_repository(
        name = "com_github_gobuffalo_genny",
        importpath = "github.com/gobuffalo/genny",
        sum = "h1:gJb0C7JcowRQ6bFIjhhwUhqXZ6d1CdBYZyliCRTLbRs=",
        version = "v0.0.0-20190315121735-8b38fb089e88",
    )
    go_repository(
        name = "com_github_gobuffalo_github_flavored_markdown",
        importpath = "github.com/gobuffalo/github_flavored_markdown",
        sum = "h1:Vjvz4wqOnviiLEfTh5bh270b3lhpJiwwQEWOWmHMwY8=",
        version = "v1.0.7",
    )
    go_repository(
        name = "com_github_gobuffalo_gogen",
        importpath = "github.com/gobuffalo/gogen",
        sum = "h1:f3Fpd5AqsFuTHUEhUeEMIFJkX8FpVnzdW+GpYxIyXkA=",
        version = "v0.0.0-20190315121717-8f38393713f5",
    )
    go_repository(
        name = "com_github_gobuffalo_httptest",
        importpath = "github.com/gobuffalo/httptest",
        sum = "h1:LWp2khlgA697h4BIYWW2aRxvB93jMnBrbakQ/r2KLzs=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_gobuffalo_licenser",
        importpath = "github.com/gobuffalo/licenser",
        sum = "h1:bbR/354ALkf/k/CoHxgPPAToVU38WZVwXcxF/UgT7V4=",
        version = "v0.0.0-20181211173111-f8a311c51159",
    )
    go_repository(
        name = "com_github_gobuffalo_logger",
        importpath = "github.com/gobuffalo/logger",
        sum = "h1:8thhT+kUJMTMy3HlX4+y9Da+BNJck+p109tqqKp7WDs=",
        version = "v0.0.0-20190315122211-86e12af44bc2",
    )
    go_repository(
        name = "com_github_gobuffalo_makr",
        importpath = "github.com/gobuffalo/makr",
        sum = "h1:lOlpv2iz0dNa4qse0ZYQgbtT+ybwVxWEAcOZbcPmeYc=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_github_gobuffalo_mapi",
        importpath = "github.com/gobuffalo/mapi",
        sum = "h1:JRuTiZzDEZhBHkFiHTxJkYRT6CbYuL0K/rn+1byJoEA=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_gobuffalo_meta",
        importpath = "github.com/gobuffalo/meta",
        sum = "h1:cP3lJDiGboBok8q6axF0rqWjlg/MOCfgjuwuNp5TlhE=",
        version = "v0.0.0-20190120163247-50bbb1fa260d",
    )
    go_repository(
        name = "com_github_gobuffalo_mw_basicauth",
        importpath = "github.com/gobuffalo/mw-basicauth",
        sum = "h1:bCqDBHnByenQitOtFdEtMvlWVgPwODrfZ+nVkgGoJZ8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_gobuffalo_mw_contenttype",
        importpath = "github.com/gobuffalo/mw-contenttype",
        sum = "h1:SUFp8EbFjlKXkvqstoxPWx3nVPV3BSKZTswQNTZFaik=",
        version = "v0.0.0-20180802152300-74f5a47f4d56",
    )
    go_repository(
        name = "com_github_gobuffalo_mw_csrf",
        importpath = "github.com/gobuffalo/mw-csrf",
        sum = "h1:A13B4mhcFQcjPJ1GFBrh61B4Qo87fZa82FfTt9LX/QU=",
        version = "v0.0.0-20180802151833-446ff26e108b",
    )
    go_repository(
        name = "com_github_gobuffalo_mw_forcessl",
        importpath = "github.com/gobuffalo/mw-forcessl",
        sum = "h1:v94+IGhlBro0Lz1gOR3lrdAVSZ0mJF2NxsdppKd7FnI=",
        version = "v0.0.0-20180802152810-73921ae7a130",
    )
    go_repository(
        name = "com_github_gobuffalo_mw_i18n",
        importpath = "github.com/gobuffalo/mw-i18n",
        sum = "h1:pZhsgF8RXEngHdibuRNOXNk1pL0K9rFa5HOcvURNTQ4=",
        version = "v0.0.0-20180802152014-e3060b7e13d6",
    )
    go_repository(
        name = "com_github_gobuffalo_mw_paramlogger",
        importpath = "github.com/gobuffalo/mw-paramlogger",
        sum = "h1:TsmUXyHjj5ReuN1AJjEVukf72J6AfRTF2CfTEaqVLT8=",
        version = "v0.0.0-20181005191442-d6ee392ec72e",
    )
    go_repository(
        name = "com_github_gobuffalo_mw_tokenauth",
        importpath = "github.com/gobuffalo/mw-tokenauth",
        sum = "h1:nhPzONHNGlXZIMFfKm6cWpRSq5oTanRK1qBtfCPBFyE=",
        version = "v0.0.0-20181001105134-8545f626c189",
    )
    go_repository(
        name = "com_github_gobuffalo_packd",
        importpath = "github.com/gobuffalo/packd",
        sum = "h1:P6naWPiHm/7R3eYx/ub3VhaW9G+1xAMJ6vzACePaGPI=",
        version = "v0.0.0-20190315124812-a385830c7fc0",
    )
    go_repository(
        name = "com_github_gobuffalo_packr",
        importpath = "github.com/gobuffalo/packr",
        sum = "h1:pnDU7VZ1gPbkIEDiXzWPtUZr/RgWSpq9x7PE7uQB8zo=",
        version = "v1.24.0",
    )
    go_repository(
        name = "com_github_gobuffalo_packr_v2",
        importpath = "github.com/gobuffalo/packr/v2",
        sum = "h1:iCzGL6YeqMNrrch5mwZpHWaUlFhzjf5G8YzsBrbAjTo=",
        version = "v2.0.6",
    )
    go_repository(
        name = "com_github_gobuffalo_plush",
        importpath = "github.com/gobuffalo/plush",
        sum = "h1:sxY0WMa6J1pMuomTUQ8n2TlR3otiCVZaq21gSrHLScU=",
        version = "v3.7.32+incompatible",
    )
    go_repository(
        name = "com_github_gobuffalo_plushgen",
        importpath = "github.com/gobuffalo/plushgen",
        sum = "h1:y+xgpBg0jHuTTDOE65FahfUgugBDpOp4C0RXXtGiSB0=",
        version = "v0.0.0-20190104222512-177cd2b872b3",
    )
    go_repository(
        name = "com_github_gobuffalo_pop",
        importpath = "github.com/gobuffalo/pop",
        sum = "h1:5v15ZgICK3MFTU90QRqCaqDUf4wcriIbws1hqpYL2Xo=",
        version = "v4.8.4+incompatible",
    )
    go_repository(
        name = "com_github_gobuffalo_release",
        importpath = "github.com/gobuffalo/release",
        sum = "h1:haad/N5NdoEKK4bmaHLuMeuZdNYIJyXmv+TA8ZwFYwo=",
        version = "v1.1.6",
    )
    go_repository(
        name = "com_github_gobuffalo_shoulders",
        importpath = "github.com/gobuffalo/shoulders",
        sum = "h1:BqVJBUXlBWAf+WLhXijVk3SCpp75LXrVBiIkOCzZbNc=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_gobuffalo_syncx",
        importpath = "github.com/gobuffalo/syncx",
        sum = "h1:tpom+2CJmpzAWj5/VEHync2rJGi+epHNIeRSWjzGA+4=",
        version = "v0.0.0-20190224160051-33c29581e754",
    )
    go_repository(
        name = "com_github_gobuffalo_tags",
        importpath = "github.com/gobuffalo/tags",
        sum = "h1:qc5hMXPsXD+zYPSlNgqzlpVW7z/+GJEVCqn245BEUIY=",
        version = "v2.0.15+incompatible",
    )
    go_repository(
        name = "com_github_gobuffalo_uuid",
        importpath = "github.com/gobuffalo/uuid",
        sum = "h1:c5uWRuEnYggYCrT9AJm0U2v1QTG7OVDAvxhj8tIV5Gc=",
        version = "v2.0.5+incompatible",
    )
    go_repository(
        name = "com_github_gobuffalo_validate",
        importpath = "github.com/gobuffalo/validate",
        sum = "h1:6f4JCEz11Zi6iIlexMv7Jz10RBPvgI795AOaubtCwTE=",
        version = "v2.0.3+incompatible",
    )
    go_repository(
        name = "com_github_gobuffalo_x",
        importpath = "github.com/gobuffalo/x",
        sum = "h1:N0iqtKwkicU8M2rLirTDJxdwuL8I2/8MjMlEayaNSgE=",
        version = "v0.0.0-20181007152206-913e47c59ca7",
    )
    go_repository(
        name = "com_github_gobwas_glob",
        importpath = "github.com/gobwas/glob",
        sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_gocql_gocql",
        importpath = "github.com/gocql/gocql",
        sum = "h1:vF83LI8tAakwEwvWZtrIEx7pOySacl2TOxx6eXk4ePo=",
        version = "v0.0.0-20190301043612-f6df8288f9b4",
    )
    go_repository(
        name = "com_github_gofrs_uuid",
        importpath = "github.com/gofrs/uuid",
        sum = "h1:y12jRkkFxsd7GpqdSZ+/KCs/fJbqpEXSGd4+jfEaewE=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:/s5zKNz0uPFCZ5hddgPdo2TK2TVrUNMn0OOX8/aZMTE=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_golang_gddo",
        importpath = "github.com/golang/gddo",
        sum = "h1:CfaPdCDbZu8jSwjq0flJv2u+WreQM0KqytUQahZ6Xf4=",
        version = "v0.0.0-20190312205958-5a2505f3dbf0",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
        version = "v0.0.0-20160126235308-23def4e6c14b",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:ZgQEtGgCBiWRM39fZuwSd1LwSqqSW0hOdXCYYDX0R3I=",
        version = "v0.0.0-20190702054246-869f871628b6",
    )
    go_repository(
        name = "com_github_golang_migrate_migrate_v4",
        importpath = "github.com/golang-migrate/migrate/v4",
        sum = "h1:gONcHxHApDTKXDyLH/H97gEHmpu1zcnnbAaq2zgrPrs=",
        version = "v4.7.0",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:qGJ6qTW+x6xX/my+8YUVl4WNpX9B7+/l2tRsHGZ7f2s=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:6nsPYzhq5kReh6QImI3k5qWzO4PEbvbIW2cwSfR/6xs=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:Qgr9rKW7uDUkrbSmQeiDsGa8SjGyCOGtuasMWwvp2P4=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:Xye71clBPdm5HgqGwUkwhbynsUJZhDbS20FvLhQ2izg=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_google_go_github",
        importpath = "github.com/google/go-github",
        sum = "h1:N0LgJ1j65A7kfXrZnUDaYCs/Sf4rEjNlfyDHW9dolSY=",
        version = "v17.0.0+incompatible",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        importpath = "github.com/google/go-querystring",
        sum = "h1:Xkwi/a1rcvNg1PPYe5vI8GbeBY/jrVuDX5ASuANWTrk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_martian",
        importpath = "github.com/google/martian",
        sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_google_pprof",
        importpath = "github.com/google/pprof",
        sum = "h1:XTnP8fJpa4Kvpw2qARB4KS9izqxPS0Sd92cDlY3uk+w=",
        version = "v0.0.0-20190723021845-34ac40c74b70",
    )
    go_repository(
        name = "com_github_google_renameio",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_subcommands",
        importpath = "github.com/google/subcommands",
        sum = "h1:/eqq+otEXm5vhfBrbREPCSVQbvofip6kIz+mX5TUH7k=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_google_wire",
        importpath = "github.com/google/wire",
        sum = "h1:imGQZGEVEHpje5056+K+cgdO72p0LQv2xIIFXNGUf60=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
        version = "v2.0.5",
    )
    go_repository(
        name = "com_github_gopherjs_gopherjs",
        importpath = "github.com/gopherjs/gopherjs",
        sum = "h1:JJ2EP5vV3LAD2U1CxQtD7PTOO15Y96kXmKDz7TjxGHs=",
        version = "v0.0.0-20181004151105-1babbf986f6f",
    )
    go_repository(
        name = "com_github_gorilla_context",
        importpath = "github.com/gorilla/context",
        sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_css",
        importpath = "github.com/gorilla/css",
        sum = "h1:BQqNyPTi50JCFMTw/b67hByjMVXZRwGha6wxVGkeihY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gorilla_handlers",
        importpath = "github.com/gorilla/handlers",
        sum = "h1:0QniY0USkHQ1RGCLfKxeNHK9bkDHGRYGNDFBCS+YARg=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sum = "h1:Dw4jY2nghMMRsh1ol8dv1axHkDwMQK2DHerMNJsIpJU=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_github_gorilla_pat",
        importpath = "github.com/gorilla/pat",
        sum = "h1:LqbZZ9sNMWVjeXS4NN5oVvhMjDyLhmA1LG86oSo+IqY=",
        version = "v0.0.0-20180118222023-199c85a7f6d1",
    )
    go_repository(
        name = "com_github_gorilla_securecookie",
        importpath = "github.com/gorilla/securecookie",
        sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_sessions",
        importpath = "github.com/gorilla/sessions",
        sum = "h1:2i3OrAjARC9UWJ/mie/5liA3dOcO9hcXIe0QHQLrTQ0=",
        version = "v1.1.4-0.20181208214519-12bd4761fc66",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:WDFjx/TMzVgy9VdMMQi2K2Emtwi2QcUQsztZ/zLaH/Q=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_gotestyourself_gotestyourself",
        importpath = "github.com/gotestyourself/gotestyourself",
        sum = "h1:AQwinXlbQR2HvPjQZOmDhRqsv5mZf+Jb1RnSLxcqZcI=",
        version = "v2.2.0+incompatible",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        sum = "h1:Iju5GlWwrvL6UBg4zJJt3btmonfrMlCDdsejg4CZE7c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:8p0pcgLlw2iuZVsdHdPaMUXFOA+6gDixcXbHEMzSyW8=",
        version = "v1.9.6",
    )
    go_repository(
        name = "com_github_gtank_cryptopasta",
        importpath = "github.com/gtank/cryptopasta",
        sum = "h1:7xsUJsB2NrdcttQPa7JLEaGzvdbk7KvfrjgHZXOQRo0=",
        version = "v0.0.0-20170601214702-1f550f6f2f69",
    )
    go_repository(
        name = "com_github_hailocab_go_hostpool",
        importpath = "github.com/hailocab/go-hostpool",
        sum = "h1:5upAirOpQc1Q53c0bnx2ufif5kANL7bfZWcc6VJWJd8=",
        version = "v0.0.0-20160125115350-e80d13ce29ed",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:YPkqC67at8FYaadspW/6uE0COsBxS2656RLEr8Bppgk=",
        version = "v0.5.3",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        importpath = "github.com/imdario/mergo",
        sum = "h1:k/Biv+LJL35wkk0Hveko1nj7as4tSHkHdZaNlzn/gcQ=",
        version = "v0.0.0-20171009183408-7fe0c75c13ab",
    )
    go_repository(
        name = "com_github_imkira_go_interpol",
        importpath = "github.com/imkira/go-interpol",
        sum = "h1:KIiKr0VSG2CUW1hl1jpiyuzuJeKUUpC8iM1AIE7N1Vk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        importpath = "github.com/inconshreveable/mousetrap",
        sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_fake",
        importpath = "github.com/jackc/fake",
        sum = "h1:vr3AYkKovP8uR8AvSGGUK1IDqRa5lAAvEkZG1LKaCRc=",
        version = "v0.0.0-20150926172116-812a484cc733",
    )
    go_repository(
        name = "com_github_jackc_pgx",
        importpath = "github.com/jackc/pgx",
        sum = "h1:0Vihzu20St42/UDsvZGdNE6jak7oi/UOeMzwMPHkgFY=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_jaytaylor_html2text",
        importpath = "github.com/jaytaylor/html2text",
        sum = "h1:jTkyeF7NZ5oIr0ESmcrpiDgAfoidCBF4F5kJhjtaRwE=",
        version = "v0.0.0-20190408195923-01ec452cbe43",
    )
    go_repository(
        name = "com_github_jessevdk_go_flags",
        importpath = "github.com/jessevdk/go-flags",
        sum = "h1:4IU2WS7AumrZ/40jfhf4QVDMsQwqA7VEHozFRrGARJA=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_jinzhu_gorm",
        importpath = "github.com/jinzhu/gorm",
        sum = "h1:gaHGvE+UnWGlbWG4Y3FUwY1EcZ5n6S9WtqBA/uySMLE=",
        version = "v1.9.11",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        importpath = "github.com/jinzhu/now",
        sum = "h1:HjfetcXq097iXP0uoPCdnM4Efp5/9MsM0/M+XOTeR3M=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:pmfjZENx5imkbgOkpRUYLnmbU7UEFbjtDA2hxJ1ichM=",
        version = "v0.0.0-20180206201540-c2b33e8439af",
    )
    go_repository(
        name = "com_github_jmoiron_sqlx",
        importpath = "github.com/jmoiron/sqlx",
        sum = "h1:41Ip0zITnmWNR/vHV+S4m+VoUivnWY5E4OJfLZjCJMA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        importpath = "github.com/joho/godotenv",
        sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_jonboulle_clockwork",
        importpath = "github.com/jonboulle/clockwork",
        sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:KfgG9LzI+pYjr4xvmz/5H4FXjokeP+rlHLhv3iH62Fo=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_jstemmer_go_junit_report",
        importpath = "github.com/jstemmer/go-junit-report",
        sum = "h1:rBMNdlhTLzJjJSDIjNEXX1Pz3Hmwmz91v+zycvx9PJc=",
        version = "v0.0.0-20190106144839-af01ea7f8024",
    )
    go_repository(
        name = "com_github_jtolds_gls",
        importpath = "github.com/jtolds/gls",
        sum = "h1:fSuqC+Gmlu6l/ZYAoZzx2pyucC8Xza35fpRVWLVmUEE=",
        version = "v4.2.1+incompatible",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_k0kubun_colorstring",
        importpath = "github.com/k0kubun/colorstring",
        sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
        version = "v0.0.0-20150214042306-9440f1994b88",
    )
    go_repository(
        name = "com_github_kardianos_osext",
        importpath = "github.com/kardianos/osext",
        sum = "h1:iQTw/8FWTuc7uiaSepXwyf3o52HaUYcV+Tu66S3F5GA=",
        version = "v0.0.0-20190222173326-2bc1f35cddc0",
    )
    go_repository(
        name = "com_github_karrick_godirwalk",
        importpath = "github.com/karrick/godirwalk",
        sum = "h1:ycpSqVon/QJJoaT1t8sae0tp1Stg21j+dyuS7OoagcA=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_kballard_go_shellquote",
        importpath = "github.com/kballard/go-shellquote",
        sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
        version = "v0.0.0-20180428030007-95032a82bc51",
    )
    go_repository(
        name = "com_github_kisielk_errcheck",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:reN85Pxc5larApoH1keMBiu2GWtPqXQ1nc9gx+jOU+E=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:Bx0qjetmNjdFXASH02NSAREKpiaDwkO1DRZ3dV2KCcs=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_klauspost_cpuid",
        importpath = "github.com/klauspost/cpuid",
        sum = "h1:vJi+O/nMdFt0vqm8NZBI6wzALWdA2X+egi0ogNyrC/w=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:DB17ag19krx9CFsz4o3enTrPXyIXCl+2iCXH/aMAp9s=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_kr_logfmt",
        importpath = "github.com/kr/logfmt",
        sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
        version = "v0.0.0-20140226030751-b84e30acd515",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:AkaSdXYQOWeaO3neb8EM634ahkXXe3jYbVh/F9lq+GI=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kshvakov_clickhouse",
        importpath = "github.com/kshvakov/clickhouse",
        sum = "h1:PDTYk9VYgbjPAWry3AoDREeMgOVUFij6bh6IjlloHL0=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_leodido_go_urn",
        importpath = "github.com/leodido/go-urn",
        sum = "h1:hpXL4XnriNwQ/ABnpepYM/1vCLWNDfUNts8dX3xTG6Y=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq",
        sum = "h1:sJZmqHoEaY7f+NPP8pgLB/WxulyR3fewgCM2qaSlBb4=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_luna_duclos_instrumentedsql",
        importpath = "github.com/luna-duclos/instrumentedsql",
        sum = "h1:KWGe9RPYdAnJMmFpP0HlnlthjfvQ0pWA9hrQlhz/cdM=",
        version = "v0.0.0-20190316074304-ecad98b20aec",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:aizVhC/NAAcKWb+5QsU1iNOZb4Yws5UO2I+aIprQITM=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_markbates_deplist",
        importpath = "github.com/markbates/deplist",
        sum = "h1:BKTJDTV5EynLGvTyONdgYVvV34DWq20mJAfGBCP+AYs=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_markbates_going",
        importpath = "github.com/markbates/going",
        sum = "h1:uNQHDDfMRNOUmuxDbPbvatyw4wr4UOSUZkGkdkcip1o=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_markbates_grift",
        importpath = "github.com/markbates/grift",
        sum = "h1:JjTyhlgPtgEnyHNvVn5lk21zWQbWD3cGE0YdyvvbZYg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_markbates_hmax",
        importpath = "github.com/markbates/hmax",
        sum = "h1:yo2N0gBoCnUMKhV/VRLHomT6Y9wUm+oQQENuWJqCdlM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_markbates_inflect",
        importpath = "github.com/markbates/inflect",
        sum = "h1:5fh1gzTFhfae06u3hzHYO9xe3l3v3nW5Pwt3naLTP5g=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_markbates_oncer",
        importpath = "github.com/markbates/oncer",
        sum = "h1:JgVTCPf0uBVcUSWpyXmGpgOc62nK5HWUBKAGc3Qqa5k=",
        version = "v0.0.0-20181203154359-bf2de49a0be2",
    )
    go_repository(
        name = "com_github_markbates_refresh",
        importpath = "github.com/markbates/refresh",
        sum = "h1:6EZ/vvVpWiam8OTIhrhfV9cVJR/NvScvcCiqosbTkbA=",
        version = "v1.4.10",
    )
    go_repository(
        name = "com_github_markbates_safe",
        importpath = "github.com/markbates/safe",
        sum = "h1:yjZkbvRM6IzKj9tlu/zMJLS0n/V351OZWRnF3QfaUxI=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_markbates_sigtx",
        importpath = "github.com/markbates/sigtx",
        sum = "h1:y/xtkBvNPRjD4KeEplf4w9rJVSc23/xl+jXYGowTwy0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_markbates_willie",
        importpath = "github.com/markbates/willie",
        sum = "h1:394PpHImWjScL9X2VRCDXJAcc77sHsSr3w3sOnL/DVc=",
        version = "v1.0.9",
    )
    go_repository(
        name = "com_github_masterminds_semver",
        importpath = "github.com/Masterminds/semver",
        sum = "h1:WBLTQ37jOCzSLtXNdoo8bNM8876KhNqOKvrlGITgsTc=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:/bC9yWikZXAL9uJdulbSfyVNIR3n3trXl+v8+1sx8mU=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:HLtExJ+uU2HOZ+wI0Tt5DtUDrx8yhUqDcp7fYERX4CE=",
        version = "v0.0.8",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:V2iyH+aX9C5fsYCpK60U8BYIvmhqxuOL3JZcqc1NB7k=",
        version = "v0.0.6",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:LDdKkqtYlom37fkvqs8rMPFKAMe8+SgjbwZ6ex1/A/Q=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_mattn_goveralls",
        importpath = "github.com/mattn/goveralls",
        sum = "h1:7eJB6EqsPhRVxvwEXGnqdO2sJI0PTsrWoTMXEk9/OQc=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_meatballhat_negroni_logrus",
        importpath = "github.com/meatballhat/negroni-logrus",
        sum = "h1:V6GHkMOIsnpGDasS1iYiNxEYTY8TmyjQXEF8PqYkKQ8=",
        version = "v0.0.0-20170801195057-31067281800f",
    )
    go_repository(
        name = "com_github_mendsley_gojwk",
        importpath = "github.com/mendsley/gojwk",
        sum = "h1:Z/i1e+gTZrmcGeZyWckaLfucYG6KYOXLWo4co8pZYNY=",
        version = "v0.0.0-20141217222730-4d5ec6e58103",
    )
    go_repository(
        name = "com_github_microcosm_cc_bluemonday",
        importpath = "github.com/microcosm-cc/bluemonday",
        sum = "h1:5lPfLTTAvAbtS0VqT+94yOtFnGfUWYyx0+iToC3Os3s=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:xAfWHN1IrQ0NJ9TBC0KBZoqLjzDTr1ML+4MywiUOryc=",
        version = "v0.4.12",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_mohae_deepcopy",
        importpath = "github.com/mohae/deepcopy",
        sum = "h1:RWengNIwukTxcDr9M+97sNutRR1RKhG96O6jWumTTnw=",
        version = "v0.0.0-20170929034955-c48cc78d4826",
    )
    go_repository(
        name = "com_github_monoculum_formam",
        importpath = "github.com/monoculum/formam",
        sum = "h1:FEJJhVHSH+Kyxa5qNe/7dprlZbFcj2TG51OWIouhwls=",
        version = "v0.0.0-20180901015400-4e68be1d79ba",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nXxl5PrvVm2L/wCy8dQu6DMTwH4oIuGN8GJDAlqDdVE=",
        version = "v0.0.0-20170113033406-39771216ff4c",
    )
    go_repository(
        name = "com_github_moul_http2curl",
        importpath = "github.com/moul/http2curl",
        sum = "h1:Pip12xNtMvEFUBF4f8/b5yRXj94LLrNdLWELfOr2KcY=",
        version = "v0.0.0-20170919181001-9ac6cf4d929b",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_nakagami_firebirdsql",
        importpath = "github.com/nakagami/firebirdsql",
        sum = "h1:P48LjvUQpTReR3TQRbxSeSBsMXzfK0uol7eRcr7VBYQ=",
        version = "v0.0.0-20190310045651-3c02a58cfed8",
    )
    go_repository(
        name = "com_github_nicksnyder_go_i18n",
        importpath = "github.com/nicksnyder/go-i18n",
        sum = "h1:5AzlPKvXBH4qBzmZ09Ua9Gipyruv6uApMcrNZdo96+Q=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_nvveen_gotty",
        importpath = "github.com/Nvveen/Gotty",
        sum = "h1:TngWCqHvy9oXAN6lEVMRuU21PR1EtLVZJmdB18Gu3Rw=",
        version = "v0.0.0-20120604004816-cd527374f1e5",
    )
    go_repository(
        name = "com_github_oklog_ulid",
        importpath = "github.com/oklog/ulid",
        sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_oleiade_reflections",
        importpath = "github.com/oleiade/reflections",
        sum = "h1:0ir4pc6v8/PJ0yw5AEtMddfXpWBXg9cnG7SgSoJuCgY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:sq53g+DWf0J6/ceFUHpQ0nAEb6WgM++fq16MZ91cS6o=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_oneofone_xxhash",
        importpath = "github.com/OneOfOne/xxhash",
        sum = "h1:zl/OfRA6nftbBK9qTohYBJ5xvw6C/oNKizR7cZGl3cI=",
        version = "v1.2.5",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:WSHQ+IS43OoUrWtD1/bbclrwK8TTH5hzp+umCiuxHgs=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:RE1xgDvH7imwFD45h+u2SgIfERHlS2yNG4DObb5BSKU=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:WzifXhOVOEOuFYOJAW6aQqW0TooG2iki3E3Ii+WN7gQ=",
        version = "v1.0.0-rc1",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:JMemWkRwHx4Zj+fVxWoMCFm/8sYGGrUVojFA6h/TRcI=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        importpath = "github.com/opencontainers/runc",
        sum = "h1:GlxAyO6x8rfZYN9Tt0Kti5a/cP41iuiO2yYT0IJGY8Y=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_openzipkin_zipkin_go",
        importpath = "github.com/openzipkin/zipkin-go",
        sum = "h1:yXiysv1CSK7Q5yjGy1710zZGnsbMUIjluWBxtLXHPBo=",
        version = "v0.1.6",
    )
    go_repository(
        name = "com_github_ory_dockertest",
        importpath = "github.com/ory/dockertest",
        sum = "h1:iLLK6SQwIhcbrG783Dghaaa3WPzGc+4Emza6EbVUUGA=",
        version = "v3.3.5+incompatible",
    )
    go_repository(
        name = "com_github_ory_fosite",
        importpath = "github.com/ory/fosite",
        sum = "h1:1HQZPXQ+PdeIe27H9Fjfvxem5uxVc/diIwUDk5XTEM4=",
        version = "v0.30.2",
    )
    go_repository(
        name = "com_github_ory_go_acc",
        importpath = "github.com/ory/go-acc",
        sum = "h1:Bpk3eqc3rbJT2mE+uS9ETzmi2cEL4RuIKz2iUeteh04=",
        version = "v0.0.0-20181118080137-ddc355013f90",
    )
    go_repository(
        name = "com_github_ory_go_convenience",
        importpath = "github.com/ory/go-convenience",
        sum = "h1:zouLKfF2GoSGnJwGq+PE/nJAE6dj2Zj5QlTgmMTsTS8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_ory_gojsonschema",
        importpath = "github.com/ory/gojsonschema",
        sum = "h1:4u2kqhTpMYabxKmTYyJasLXFozkGVvbb6+AdMnWadwg=",
        version = "v0.0.0-20190717132251-f184856edacf",
    )
    go_repository(
        name = "com_github_ory_graceful",
        importpath = "github.com/ory/graceful",
        sum = "h1:zx+8tDObLPrG+7Tc8jKYlXsqWnLtOQA1IZ/FAAKHMXU=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_ory_herodot",
        importpath = "github.com/ory/herodot",
        sum = "h1:Dcs4yH1Qw1GIgGCvvvdafhT8xjwElTE//8xLmHtPEYA=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_ory_hydra",
        importpath = "github.com/ory/hydra",
        sum = "h1:u6WdJ0lheQ6zhMGpmXl/wSQcEW0zXWsuFTCmyaV/SmI=",
        version = "v1.0.9",
    )
    go_repository(
        name = "com_github_ory_viper",
        importpath = "github.com/ory/viper",
        sum = "h1:w4ceGgWwWLzAFYQ7bHaDZmwNsAto2JPVdyQjQnn7VWI=",
        version = "v1.5.6",
    )
    go_repository(
        name = "com_github_ory_x",
        importpath = "github.com/ory/x",
        sum = "h1:fThIGB08t5Fyqa2sD8I6c+gE+Z3G6hJGnzk/O681Mp4=",
        version = "v0.0.73",
    )
    go_repository(
        name = "com_github_oxtoacart_bpool",
        importpath = "github.com/oxtoacart/bpool",
        sum = "h1:rp5dCmg/yLR3mgFuSOe4oEnDDmGLROTvMragMUXpTQw=",
        version = "v0.0.0-20190530202638-03653db5a59c",
    )
    go_repository(
        name = "com_github_parnurzeal_gorequest",
        importpath = "github.com/parnurzeal/gorequest",
        sum = "h1:oPjDCsF5IkD4gUk6vIgsxYNaSgvAnIh1EJeROn3HdJU=",
        version = "v0.2.15",
    )
    go_repository(
        name = "com_github_pborman_uuid",
        importpath = "github.com/pborman/uuid",
        sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_pelletier_go_toml",
        importpath = "github.com/pelletier/go-toml",
        sum = "h1:aetoXYr0Tv7xRU/V4B4IZJ2QcbtMUFoNb3ORp7TzIK4=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_phayes_freeport",
        importpath = "github.com/phayes/freeport",
        sum = "h1:rZQtoozkfsiNs36c7Tdv/gyGNzD1X1XWKO8rptVNZuM=",
        version = "v0.0.0-20171002181615-b8543db493a5",
    )
    go_repository(
        name = "com_github_pierrec_lz4",
        importpath = "github.com/pierrec/lz4",
        sum = "h1:2xWsjqPFWcplujydGg4WmhC/6fZqK42wMM8aXeqhl0I=",
        version = "v2.0.5+incompatible",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:iURUrRGxPUNPdy5/HRSm+Yj6okJ6UtLINN0Q9M4+h3I=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_github_pkg_profile",
        importpath = "github.com/pkg/profile",
        sum = "h1:OQIvuDgm00gWVWGTf4m4mCt6W1/0YqU7Ntg0mySWgaI=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_pquerna_cachecontrol",
        importpath = "github.com/pquerna/cachecontrol",
        sum = "h1:J9b7z+QKAmPf4YLrFg6oQUotqHQeUNWwkvo7jZp1GLU=",
        version = "v0.0.0-20180517163645-1555304b9b35",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:BQ53HtBmfOitExawJ6LokA4x8ov/z0SYYb0+HxJfRI8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version = "v0.0.0-20190812154241-14fe0d1b01d4",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:kRhiuYSXR3+uv2IbVbZhUxK5zVD/2pp3Gd2PpvPkpEo=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:CTwfnzjQ+8dS6MhHHu4YswVAD99sL2wjPqP+VkURmKE=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_prometheus_tsdb",
        importpath = "github.com/prometheus/tsdb",
        sum = "h1:If5rVCMTp6W2SiRAQFlbpJNgVlgMEd+U2GZckwK38ic=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_puerkitobio_goquery",
        importpath = "github.com/PuerkitoBio/goquery",
        sum = "h1:uGvmFXOA73IKluu/F84Xd1tt/z07GYm8X49XKHP7EJk=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_puerkitobio_purell",
        importpath = "github.com/PuerkitoBio/purell",
        sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_puerkitobio_urlesc",
        importpath = "github.com/PuerkitoBio/urlesc",
        sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
        version = "v0.0.0-20170810143723-de5bf2ad4578",
    )
    go_repository(
        name = "com_github_rcrowley_go_metrics",
        importpath = "github.com/rcrowley/go-metrics",
        sum = "h1:9ZKAASQSHhDYGoxY8uLVpewe1GDZ2vu2Tr/vTdVAkFQ=",
        version = "v0.0.0-20181016184325-3113b8401b8a",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:RR9dF3JtopPvtkroDZuVD7qquD0bnHlKSqaQhgwt8yk=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_rs_cors",
        importpath = "github.com/rs/cors",
        sum = "h1:G9tHG9lebljV9mfp9SNPDL36nCDxmo3zTlAf1YgvzmI=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_rs_xid",
        importpath = "github.com/rs/xid",
        sum = "h1:mhH9Nq+C1fY2l1XIpgxIiUOfNpRBYH1kKcr+qfKgjRc=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        importpath = "github.com/rs/zerolog",
        sum = "h1:RMRHFw2+wF7LO0QqtELQwo8hqSmqISyCJeFeAAuWcRo=",
        version = "v1.17.2",
    )
    go_repository(
        name = "com_github_rubenv_sql_migrate",
        importpath = "github.com/rubenv/sql-migrate",
        sum = "h1:ID2fzWzRFJcF/xf/8eLN9GW5CXb6NQnKfC+ksTwMNpY=",
        version = "v0.0.0-20190212093014-1007f53448d7",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_satori_go_uuid",
        importpath = "github.com/satori/go.uuid",
        sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_sawadashota_encrypta",
        importpath = "github.com/sawadashota/encrypta",
        sum = "h1:R46/RxYmYdxI3VOt63B637OVBHzu+fazPyLo5CqK6QE=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_segmentio_analytics_go",
        importpath = "github.com/segmentio/analytics-go",
        sum = "h1:W7T3ieNQjPFMb+SE8SAVYo6mPkKK/Y37wYdiNf5lCVg=",
        version = "v3.0.1+incompatible",
    )
    go_repository(
        name = "com_github_segmentio_backo_go",
        importpath = "github.com/segmentio/backo-go",
        sum = "h1:rsRTAcCR5CeNLkvgBVSjQoDGRRt6kggsE6XYBqCv2KQ=",
        version = "v0.0.0-20160424052352-204274ad699c",
    )
    go_repository(
        name = "com_github_serenize_snaker",
        importpath = "github.com/serenize/snaker",
        sum = "h1:ofR1ZdrNSkiWcMsRrubK9tb2/SlZVWttAfqUjJi6QYc=",
        version = "v0.0.0-20171204205717-a683aaf2d516",
    )
    go_repository(
        name = "com_github_sergi_go_diff",
        importpath = "github.com/sergi/go-diff",
        sum = "h1:Kpca3qRNrduNnOQeazBd0ysaKrUJiIuISHxogkT9RPQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_shopify_sarama",
        importpath = "github.com/Shopify/sarama",
        sum = "h1:9oksLxC6uxVPHPVYUmq6xhr1BOF/hHobWH2UzO67z1s=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_github_shopify_toxiproxy",
        importpath = "github.com/Shopify/toxiproxy",
        sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
        version = "v2.1.4+incompatible",
    )
    go_repository(
        name = "com_github_shopspring_decimal",
        importpath = "github.com/shopspring/decimal",
        sum = "h1:pntxY8Ary0t43dCZ5dqY4YTJCObLY1kIXl0uzMv+7DE=",
        version = "v0.0.0-20180709203117-cd690d0c9e24",
    )
    go_repository(
        name = "com_github_shurcool_go",
        importpath = "github.com/shurcooL/go",
        sum = "h1:MZM7FHLqUHYI0Y/mQAt3d2aYa0SiNms/hFqC9qJYolM=",
        version = "v0.0.0-20180423040247-9e1955d9fb6e",
    )
    go_repository(
        name = "com_github_shurcool_go_goon",
        importpath = "github.com/shurcooL/go-goon",
        sum = "h1:llrF3Fs4018ePo4+G/HV/uQUqEI1HMDjCeOf2V6puPc=",
        version = "v0.0.0-20170922171312-37c2f522c041",
    )
    go_repository(
        name = "com_github_shurcool_highlight_diff",
        importpath = "github.com/shurcooL/highlight_diff",
        sum = "h1:vYEG87HxbU6dXj5npkeulCS96Dtz5xg3jcfCgpcvbIw=",
        version = "v0.0.0-20170515013008-09bb4053de1b",
    )
    go_repository(
        name = "com_github_shurcool_highlight_go",
        importpath = "github.com/shurcooL/highlight_go",
        sum = "h1:xLQlo0Ghg8zBaQi+tjpK+z/WLjbg/BhAWP9pYgqo/LQ=",
        version = "v0.0.0-20170515013102-78fb10f4a5f8",
    )
    go_repository(
        name = "com_github_shurcool_octicon",
        importpath = "github.com/shurcooL/octicon",
        sum = "h1:j3cAp1j8k/tSLaCcDiXIpVJ8FzSJ9g1eeOAPRJYM75k=",
        version = "v0.0.0-20180602230221-c42b0e3b24d9",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        sum = "h1:/vdW8Cb7EXrkqWGufVMES1OH2sU9gKVb2n9/1y5NMBY=",
        version = "v0.0.0-20170918181015-86672fcb3f95",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:SPIRibHv4MatM3XXNO2BJeFLZwZ2LvZgfQ5+UNI2im4=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_smartystreets_assertions",
        importpath = "github.com/smartystreets/assertions",
        sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
        version = "v0.0.0-20180927180507-b2de0cb4f26d",
    )
    go_repository(
        name = "com_github_smartystreets_goconvey",
        importpath = "github.com/smartystreets/goconvey",
        sum = "h1:JSvGDIbmil4Ui/dDdFBExb7/cmkNjyX5F97oglmvCDo=",
        version = "v0.0.0-20180222194500-ef6db91d284a",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        importpath = "github.com/soheilhy/cmux",
        sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_sourcegraph_annotate",
        importpath = "github.com/sourcegraph/annotate",
        sum = "h1:yKm7XZV6j9Ev6lojP2XaIshpT4ymkqhMeSghO5Ps00E=",
        version = "v0.0.0-20160123013949-f4cad6c6324d",
    )
    go_repository(
        name = "com_github_sourcegraph_syntaxhighlight",
        importpath = "github.com/sourcegraph/syntaxhighlight",
        sum = "h1:qpG93cPwA5f7s/ZPBJnGOYQNK/vKsaDaseuKT5Asee8=",
        version = "v0.0.0-20170531221838-bd320f5d308e",
    )
    go_repository(
        name = "com_github_spaolacci_murmur3",
        importpath = "github.com/spaolacci/murmur3",
        sum = "h1:7c1g84S4BPRrfL5Xrdp6fOJ206sU9y293DDHaoy0bLI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        sum = "h1:f0B+LkLX6DtmRH1isoNA9VTtNUK9K8xYd28JNNfOv/s=",
        version = "v0.0.5",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        sum = "h1:GpsTwfsQ27oS/Aha/6d1oD7tpKIqWnOA6tgOX9HHkt4=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_sqs_goreturns",
        importpath = "github.com/sqs/goreturns",
        sum = "h1:iD+PFTQwKEmbwSdwfvP5ld2WEI/g7qbdhmHJ2ASfYGs=",
        version = "v0.0.0-20181028201513-538ac6014518",
    )
    go_repository(
        name = "com_github_ssor_bom",
        importpath = "github.com/ssor/bom",
        sum = "h1:pvbZ0lM0XWPBqUKqFU8cmavspvIl9nulOYwdy6IFRRo=",
        version = "v0.0.0-20170718123548-6386211fdfcf",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:Hbg2NidpLE8veEBkEZTL3CvlkUIVzuU9jDplZO54c48=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_tidwall_pretty",
        importpath = "github.com/tidwall/pretty",
        sum = "h1:HsD+QiTn7sK6flMKIvNmpqz1qrpP3Ps6jOKIKMooyg4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        sum = "h1:LnC5Kc/wtumK+WB441p7ynQJzVuNRJiqddSIE3IlSEQ=",
        version = "v0.0.0-20190109142713-0ad062ec5ee5",
    )
    go_repository(
        name = "com_github_toqueteos_webbrowser",
        importpath = "github.com/toqueteos/webbrowser",
        sum = "h1:tVP/gpK69Fx+qMJKsLE7TD8LuGWPnEV71wBN9rrstGQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_uber_go_atomic",
        importpath = "github.com/uber-go/atomic",
        sum = "h1:Azu9lPBWRNKzYXSIwRfgRuDuS0YKsK4NFhiQv98gkxo=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_uber_jaeger_client_go",
        importpath = "github.com/uber/jaeger-client-go",
        sum = "h1:Q2Pp6v3QYiocMxomCaJuwQGFt7E53bPYqEgug/AoBtY=",
        version = "v2.16.0+incompatible",
    )
    go_repository(
        name = "com_github_uber_jaeger_lib",
        importpath = "github.com/uber/jaeger-lib",
        sum = "h1:iMSCV0rmXEogjNWPh2D0xk9YVKvrtGoHJNe9ebLu/pw=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_ugorji_go",
        importpath = "github.com/ugorji/go",
        sum = "h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_ugorji_go_codec",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:2SvQaVZ1ouYrrKKwoSk2pzd4A9evlKJb9oTL+OaLUSs=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_unrolled_secure",
        importpath = "github.com/unrolled/secure",
        sum = "h1:ltz/eIXkYWdMCZbu3Rb+bUmWVTm5AqM0QM8o0uKir4U=",
        version = "v0.0.0-20181005190816-ff9db2ff917f",
    )
    go_repository(
        name = "com_github_urfave_negroni",
        importpath = "github.com/urfave/negroni",
        sum = "h1:kIimOitoypq34K7TG7DUaJ9kq/N4Ofuwi1sjz0KipXc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_bytebufferpool",
        importpath = "github.com/valyala/bytebufferpool",
        sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_fasthttp",
        importpath = "github.com/valyala/fasthttp",
        sum = "h1:uWF8lgKmeaIewWVPwi4GRq2P6+R46IgYZdxWtM+GtEY=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_valyala_tcplisten",
        importpath = "github.com/valyala/tcplisten",
        sum = "h1:0R4NLDRDZX6JcmhJgXi5E4b8Wg84ihbmUKp/GvSPEzc=",
        version = "v0.0.0-20161114210144-ceec8f93295a",
    )
    go_repository(
        name = "com_github_vanng822_css",
        importpath = "github.com/vanng822/css",
        sum = "h1:L0rPdfzq43+NV8rfIx2kA4iSSLRj2jN5ijYHoeXRwvQ=",
        version = "v0.0.0-20190504095207-a21e860bcd04",
    )
    go_repository(
        name = "com_github_vanng822_go_premailer",
        importpath = "github.com/vanng822/go-premailer",
        sum = "h1:BJcKf0TN6EK/VIDoelbhCV45MeMVdDuS1dkxh0j8xAo=",
        version = "v0.0.0-20190504065308-ac25997d78c7",
    )
    go_repository(
        name = "com_github_vektah_gqlparser",
        importpath = "github.com/vektah/gqlparser",
        sum = "h1:ZsyLGn7/7jDNI+y4SEhI4yAxRChlv15pUHMjijT+e68=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_xanzy_go_gitlab",
        importpath = "github.com/xanzy/go-gitlab",
        sum = "h1:rWtwKTgEnXyNUGrOArN7yyc3THRkpYcKXIXia9abywQ=",
        version = "v0.15.0",
    )
    go_repository(
        name = "com_github_xdg_scram",
        importpath = "github.com/xdg/scram",
        sum = "h1:u40Z8hqBAAQyv+vATcGgV0YCnDjqSL7/q/JyPhhJSPk=",
        version = "v0.0.0-20180814205039-7eeb5667e42c",
    )
    go_repository(
        name = "com_github_xdg_stringprep",
        importpath = "github.com/xdg/stringprep",
        sum = "h1:d9X0esnoa3dFsV0FG35rAT0RIhYFlPq7MiP+DW89La0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
        version = "v0.0.0-20180127040702-4e3ac2762d5f",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xiang90_probing",
        importpath = "github.com/xiang90/probing",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
        version = "v0.0.0-20190116061207-43a291ad63a2",
    )
    go_repository(
        name = "com_github_xordataexchange_crypt",
        importpath = "github.com/xordataexchange/crypt",
        sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
        version = "v0.0.3-0.20170626215501-b2862e3d0a77",
    )
    go_repository(
        name = "com_github_xtgo_uuid",
        importpath = "github.com/xtgo/uuid",
        sum = "h1:3lbZUMbMiGUW/LMkfsEABsc5zNT9+b1CvsJx47JzJ8g=",
        version = "v0.0.0-20140804021211-a0b114877d4c",
    )
    go_repository(
        name = "com_github_yalp_jsonpath",
        importpath = "github.com/yalp/jsonpath",
        sum = "h1:6fRhSjgLCkTD3JnJxvaJ4Sj+TYblw757bqYgZaOq5ZY=",
        version = "v0.0.0-20180802001716-5cc68e5049a0",
    )
    go_repository(
        name = "com_github_yudai_gojsondiff",
        importpath = "github.com/yudai/gojsondiff",
        sum = "h1:27cbfqXLVEJ1o8I6v3y9lg8Ydm53EKqHXAOMxEGlCOA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_yudai_golcs",
        importpath = "github.com/yudai/golcs",
        sum = "h1:BHyfKlQyqbsFN5p3IfnEUduWvb9is428/nNb5L3U01M=",
        version = "v0.0.0-20170316035057-ecda9a501e82",
    )
    go_repository(
        name = "com_github_yudai_pp",
        importpath = "github.com/yudai/pp",
        sum = "h1:Q4//iY4pNF6yPLZIigmvcl7k/bPgrcTPIFIcmawg5bI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_zenazn_goji",
        importpath = "github.com/zenazn/goji",
        sum = "h1:RSQQAbXGArQ0dIDEq+PI6WqN6if+5KHu6x2Cx/GXLTQ=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_ziutek_mymysql",
        importpath = "github.com/ziutek/mymysql",
        sum = "h1:GB0qdRGsTwQSBVYuVShFBKaXSnSnYYC2d9knnE1LHFs=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_gitlab_nyarla_go_crypt",
        importpath = "gitlab.com/nyarla/go-crypt",
        sum = "h1:7gd+rd8P3bqcn/96gOZa3F5dpJr/vEiDQYlNb/y2uNs=",
        version = "v0.0.0-20160106005555-d9a5dc2b789b",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:0sMegbmn/8uTwpNkB0q9cLEpZ2W5a6kl+wtBQgPWBJQ=",
        version = "v0.44.3",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:hL+ycaJpVE9M7nLoiXb/Pn10ENE2u+oddxbD8uu0ZVU=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:Kt+gOPPp2LEPWp8CSfxhsM8ik9CcyE/gYu+0r+RnZvM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "in_gopkg_airbrake_gobrake_v2",
        importpath = "gopkg.in/airbrake/gobrake.v2",
        sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
        version = "v2.0.9",
    )
    go_repository(
        name = "in_gopkg_alecthomas_kingpin_v2",
        importpath = "gopkg.in/alecthomas/kingpin.v2",
        sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
        version = "v2.2.6",
    )
    go_repository(
        name = "in_gopkg_alexcesaro_quotedprintable_v3",
        importpath = "gopkg.in/alexcesaro/quotedprintable.v3",
        sum = "h1:2gGKlE2+asNV9m7xrywl36YYNnBG5ZQ0r/BOOxqPpmk=",
        version = "v3.0.0-20150716171945-2caba252f4dc",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:qIbj1fsPNlZgppZ+VLlY7N33q108Sa+fhmuc+sWQYwY=",
        version = "v1.0.0-20180628173108-788fd7840127",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
        importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
        sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
        version = "v2.1.2",
    )
    go_repository(
        name = "in_gopkg_go_playground_assert_v1",
        importpath = "gopkg.in/go-playground/assert.v1",
        sum = "h1:xoYuJVE7KT85PYWrN730RguIQO0ePzVRfFMXadIrXTM=",
        version = "v1.2.1",
    )
    go_repository(
        name = "in_gopkg_go_playground_validator_v9",
        importpath = "gopkg.in/go-playground/validator.v9",
        sum = "h1:icxYLlYflpazIV3ufMoNB9h9SYMQ37DZ8CTwkU4pnOs=",
        version = "v9.30.2",
    )
    go_repository(
        name = "in_gopkg_gomail_v2",
        importpath = "gopkg.in/gomail.v2",
        sum = "h1:n7WqCuqOuCbNr617RXOY0AWRXxgwEyPp2z+p0+hgMuE=",
        version = "v2.0.0-20160411212932-81ebce5c23df",
    )
    go_repository(
        name = "in_gopkg_gorp_v1",
        importpath = "gopkg.in/gorp.v1",
        sum = "h1:j3DWlAyGVv8whO7AcIWznQ2Yj7yJkn34B8s63GViAAw=",
        version = "v1.7.2",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )
    go_repository(
        name = "in_gopkg_mail_v2",
        importpath = "gopkg.in/mail.v2",
        sum = "h1:a3llQg4+Czqaf+QH4diHuHiKv4j1abMwuRXwaRNHTPU=",
        version = "v2.0.0-20180731213649-a0242b2233b4",
    )
    go_repository(
        name = "in_gopkg_resty_v1",
        importpath = "gopkg.in/resty.v1",
        sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
        version = "v1.12.0",
    )
    go_repository(
        name = "in_gopkg_square_go_jose_v2",
        importpath = "gopkg.in/square/go-jose.v2",
        sum = "h1:SK5KegNXmKmqE342YYN2qPHEnUYeoMiXXl1poUlI+o4=",
        version = "v2.3.1",
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:VUgggvou5XRW9mHwD/yXxIYSMtY0zoKQf/v226p2nyo=",
        version = "v2.2.7",
    )
    go_repository(
        name = "io_etcd_go_bbolt",
        importpath = "go.etcd.io/bbolt",
        sum = "h1:MUGmc65QhB3pIlaQ5bB4LwqSj6GIonVJXpZiaKNyaKk=",
        version = "v1.3.3",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:C9hSCOW830chIVkdja34wa6Ky+IzWllkUinR+BtRZd4=",
        version = "v0.22.0",
    )
    go_repository(
        name = "io_rsc_binaryregexp",
        importpath = "rsc.io/binaryregexp",
        sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:VGGbLNyPF7dvYHhcUGYBBGCRDDK0RRJAI6KCvo0CL+E=",
        version = "v0.8.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:QzqyMA1tlu6CgqCDUtU9V+ZKhLFT2dkJuANu5QaxI3I=",
        version = "v1.6.1",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:P/BlPoYr1gpKHOHL0/Opzbiu5X5yb55Ef4P/YGrRwno=",
        version = "v0.0.0-20190817000702-55e96fffbd48",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:AzbTB6ux+okLTzP8Ru1Xs41C303zdcfEht7MQnYJt5A=",
        version = "v1.23.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:9bFeDpN3gTqNanMVqNcoR/pJQuP5uroC3t1D7eXozTE=",
        version = "v0.0.0-20191119213627-4f8c1d86b1ba",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:estk1glOnSVeJ9tdEZZc5mAMDZk5lNJNyJ6DvrBkTEU=",
        version = "v0.0.0-20190731235908-ec7cb31e5a56",
    )
    go_repository(
        name = "org_golang_x_image",
        importpath = "golang.org/x/image",
        sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
        version = "v0.0.0-20190802002840-cff245a6509b",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:QzoH/1pFpZguR8NrRHLcO6jKqfv2zpuSqZLgdm7ZmjI=",
        version = "v0.0.0-20190409202823-959b441ac422",
    )
    go_repository(
        name = "org_golang_x_mobile",
        importpath = "golang.org/x/mobile",
        sum = "h1:+jOHrd35hpMwS0Q9obZ1Gy3kFgkSYuXO4Dh3Ibk78lY=",
        version = "v0.0.0-20190814143026-e8b3e6111d02",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:sfUMP1Gu8qASkorDVjnMuvgJzwFbTZSeXFiGBYAVdl4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:MlY3mEfbnWGmUi4rtHOtNnnnN4UJRGSyLPx+DXA5Sq4=",
        version = "v0.0.0-20191119073136-fc4aabc6c914",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:SVwTIAaPC2U/AvvLNZ2a7OVsmBpC8L5BlwK1whH3hm0=",
        version = "v0.0.0-20190604053449-0f29369cfe45",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:8gQV6CLnAEikrhgkHFbMAEhagSSnXWGV915qUMm9mrU=",
        version = "v0.0.0-20190423024810-112230192c58",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:N7DeIrjYszNmSW409R3frPPwglRwMkXSBzwVbkOjLLA=",
        version = "v0.0.0-20191120155948-bd437916bb0e",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=",
        version = "v0.3.2",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:SvFZT6jyqRaOeXpc5h/JSfZenJ2O330aBsf7JfSUXmQ=",
        version = "v0.0.0-20190308202827-9d24e82272b4",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:4cFkmztxtMslUX2SctSl+blCyXfpzhGOy9LhKAqSMA4=",
        version = "v0.0.0-20190828213141-aed303cbaa74",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:9zdDQZ7Thm29KFXgAX/+yaf3eVbP7djjWp/dXAppNCc=",
        version = "v0.0.0-20190717185122-a985d3407aa7",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:++7u8r9adKhGR+I79NfEtYrk2ktjenErXM99PSufIoI=",
        version = "v1.1.3",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:cxzIVoETapQEqDhQu3QfnvXAV4AlzcvUCxkVUFw3+EU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:HoEmRHQPVSqub6w2z2d2EOVs2fjyFRGyofhKuyDq0QI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:ORx85nbTijNz8ljznvCMR1ZBIPKFn3jQrag10X2AsuM=",
        version = "v1.10.0",
    )
    go_repository(
        name = "tools_gotest",
        importpath = "gotest.tools",
        sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
        version = "v2.2.0+incompatible",
    )
