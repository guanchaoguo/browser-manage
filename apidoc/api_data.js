define({ "api": [
  {
    "type": "Delete",
    "url": "/deleteAccount/{id}",
    "title": "删除账号",
    "description": "<p>删除账号</p>",
    "group": "account",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/account.go",
    "groupTitle": "account",
    "name": "DeleteDeleteaccountId"
  },
  {
    "type": "Get",
    "url": "/accountList",
    "title": "账号列表",
    "description": "<p>账号列表</p>",
    "group": "account",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "search",
            "description": "<p>输入框条件</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "per_page",
            "description": "<p>每页显示数据条数，默认15条</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>当前的所在页码，默认第1页</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"current_page\": 1,\n\t\"data\": [\n\t  {\n\t\t\"id\": 1,\n\t\t\"user_name\": \"admin\",\n\t\t\"password\": \"fd3c1140e2660132dc39e8788884cc8d\",\n\t\t\"salt\": \"123456\",\n\t\t\"account\": \"开发用户\",\n\t\t\"last_date\": \"0000-00-00 00:00:00\",\n\t\t\"last_login_date\": \"2017-11-07 16:13:17\",\n\t\t\"last_login_ip\": \"127.0.0.1\"\n\t  },\n\t],\n\t\"last_page\": 1,\n\t\"per_page\": 15,\n\t\"total\": 7\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/account.go",
    "groupTitle": "account",
    "name": "GetAccountlist"
  },
  {
    "type": "Get",
    "url": "/showAccount/{id}",
    "title": "修改时获取账号信息",
    "description": "<p>修改时获取账号信息</p>",
    "group": "account",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 400,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"data\": {\n\t  \"account\": \"黄飞鸿\",\n\t  \"id\": 2,\n\t  \"user_name\": \"huang\"\n\t}\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/account.go",
    "groupTitle": "account",
    "name": "GetShowaccountId"
  },
  {
    "type": "Post",
    "url": "/addAccount",
    "title": "添加账号",
    "description": "<p>添加账号</p>",
    "group": "account",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user_name",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "account",
            "description": "<p>真实姓名</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/account.go",
    "groupTitle": "account",
    "name": "PostAddaccount"
  },
  {
    "type": "Put",
    "url": "/saveAccount/{id}",
    "title": "修改保存账号",
    "description": "<p>修改保存账号</p>",
    "group": "account",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user_name",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "account",
            "description": "<p>真实姓名</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/account.go",
    "groupTitle": "account",
    "name": "PutSaveaccountId"
  },
  {
    "type": "Put",
    "url": "/savePwd/{id}",
    "title": "修改账户密码",
    "description": "<p>修改账户密码</p>",
    "group": "account",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/account.go",
    "groupTitle": "account",
    "name": "PutSavepwdId"
  },
  {
    "type": "Get",
    "url": "/getCaptcha",
    "title": "获取验证码",
    "description": "<p>获取验证码</p>",
    "group": "captcha",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"url\": \"127.0.0.1:8080/captcha/njp674vTtzlvolXVptyU7qA5zzqNiiW2\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/captcha.go",
    "groupTitle": "captcha",
    "name": "GetGetcaptcha"
  },
  {
    "type": "Post",
    "url": "/login",
    "title": "后台登录",
    "description": "<p>后台登录</p>",
    "group": "login",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user_name",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "code",
            "description": "<p>验证码</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "captchaKey",
            "description": "<p>验证码key</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 400,\n  \"message\": \"登录成功\",\n  \"result\": {\n\t\"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50Ijoi5byA5Y-R55So5oi3IiwiZGF0ZSI6MTUxMDY0NjYyMSwiaWQiOjEsInVzZXJfbmFtZSI6ImFkbWluIn0.Nd1AuCJyD0CgDPkjp3lljxhyDCBatBMrcCO-lCnE6GE\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/login.go",
    "groupTitle": "login",
    "name": "PostLogin"
  },
  {
    "type": "Post",
    "url": "/loginOut",
    "title": "退出登录",
    "description": "<p>退出登录</p>",
    "group": "login",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/login.go",
    "groupTitle": "login",
    "name": "PostLoginout"
  },
  {
    "type": "Delete",
    "url": "/deleteMenus/{id}",
    "title": "删除菜单",
    "description": "<p>删除菜单</p>",
    "group": "menus",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/menus.go",
    "groupTitle": "menus",
    "name": "DeleteDeletemenusId"
  },
  {
    "type": "Get",
    "url": "/getMenus/{id}",
    "title": "修改时获取菜单信息",
    "description": "<p>修改时获取菜单信息</p>",
    "group": "menus",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"id\": 112,\n\t\"parent_id\": 0,\n\t\"title_cn\": \"系统管理\",\n\t\"title_en\": \"\",\n\t\"class\": 0,\n\t\"desc\": \"\",\n\t\"link_url\": \"/systemManage\",\n\t\"icon\": \"icon-xitongguanli\",\n\t\"state\": 0,\n\t\"sort_id\": 0,\n\t\"menu_code\": \"M4001\",\n\t\"update_date\": \"2017-11-14 10:09:24\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/menus.go",
    "groupTitle": "menus",
    "name": "GetGetmenusId"
  },
  {
    "type": "Get",
    "url": "/menusList",
    "title": "菜单列表",
    "description": "<p>菜单列表</p>",
    "group": "menus",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"data\": [\n\t  {\n\t\t\"id\": 1,\n\t\t\"parent_id\": 0,\n\t\t\"title_cn\": \"abc\",\n\t\t\"title_en\": \"\",\n\t\t\"class\": 0,\n\t\t\"desc\": \"\",\n\t\t\"link_url\": \"\",\n\t\t\"icon\": \"\",\n\t\t\"state\": 1,\n\t\t\"sort_id\": 1,\n\t\t\"menu_code\": \"\",\n\t\t\"_child\": [\n\t\t  {\n\t\t\t\"id\": 2,\n\t\t\t\"parent_id\": 1,\n\t\t\t\"title_cn\": \"cba\",\n\t\t\t\"title_en\": \"\",\n\t\t\t\"class\": 0,\n\t\t\t\"desc\": \"\",\n\t\t\t\"link_url\": \"\",\n\t\t\t\"icon\": \"\",\n\t\t\t\"state\": 1,\n\t\t\t\"sort_id\": 1,\n\t\t\t\"menu_code\": \"\",\n\t\t\t\"_child\": [\n\t\t\t  {\n\t\t\t\t\"id\": 3,\n\t\t\t\t\"parent_id\": 2,\n\t\t\t\t\"title_cn\": \"bnm\",\n\t\t\t\t\"title_en\": \"\",\n\t\t\t\t\"class\": 0,\n\t\t\t\t\"desc\": \"\",\n\t\t\t\t\"link_url\": \"\",\n\t\t\t\t\"icon\": \"\",\n\t\t\t\t\"state\": 1,\n\t\t\t\t\"sort_id\": 1,\n\t\t\t\t\"menu_code\": \"\",\n\t\t\t\t\"_child\": []\n\t\t\t  }\n\t\t\t]\n\t\t  }\n\t\t]\n\t  }\n\t]\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/menus.go",
    "groupTitle": "menus",
    "name": "GetMenuslist"
  },
  {
    "type": "Post",
    "url": "/addMenus",
    "title": "添加菜单",
    "description": "<p>添加菜单</p>",
    "group": "menus",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title_cn",
            "description": "<p>菜单名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "link_url",
            "description": "<p>菜单链接</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "menu_code",
            "description": "<p>菜单标识符</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "icon",
            "description": "<p>菜单图标</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "sort_id",
            "description": "<p>排序 （越小越前）</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "parent_id",
            "description": "<p>父级菜单ID</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/menus.go",
    "groupTitle": "menus",
    "name": "PostAddmenus"
  },
  {
    "type": "Put",
    "url": "/saveMenus/{id}",
    "title": "修改保存菜单",
    "description": "<p>修改保存菜单</p>",
    "group": "menus",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title_cn",
            "description": "<p>菜单名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "link_url",
            "description": "<p>菜单链接</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "menu_code",
            "description": "<p>菜单标识符</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "icon",
            "description": "<p>菜单图标</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "parent_id",
            "description": "<p>父级菜单ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "sort_id",
            "description": "<p>排序 （越小越前）</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/menus.go",
    "groupTitle": "menus",
    "name": "PutSavemenusId"
  },
  {
    "type": "Delete",
    "url": "/deleteProxy/{id}",
    "title": "删除代理服务器",
    "description": "<p>删除代理服务器</p>",
    "group": "proxy",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/proxy.go",
    "groupTitle": "proxy",
    "name": "DeleteDeleteproxyId"
  },
  {
    "type": "Get",
    "url": "/proxy/{id}",
    "title": "修改时获取代理服务器信息",
    "description": "<p>修改时获取代理服务器信息</p>",
    "group": "proxy",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"id\": 2,\n\t\"addr\": \"192.168.31.155\",\n\t\"port\": 8080,\n\t\"user_name\": \"sunwukong\",\n\t\"password\": \"123456\",\n\t\"last_update_date\": \"0000-00-00 00:00:00\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/proxy.go",
    "groupTitle": "proxy",
    "name": "GetProxyId"
  },
  {
    "type": "Get",
    "url": "/proxyList",
    "title": "代理服务器列表",
    "description": "<p>代理服务器列表</p>",
    "group": "proxy",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "search",
            "description": "<p>输入框条件</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "per_page",
            "description": "<p>每页显示数据条数，默认15条</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>当前的所在页码，默认第1页</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"current_page\": 1,\n\t\"data\": [\n\t {\n\t\t\"id\": 2,\n\t\t\"addr\": \"192.168.31.155\",\n\t\t\"port\": 8080,\n\t\t\"user_name\": \"sunwukong\",\n\t\t\"password\": \"123456\",\n\t\t\"last_update_date\": \"0000-00-00 00:00:00\"\n\t  }\n\t],\n\t\"last_page\": 1,\n\t\"per_page\": 15,\n\t\"total\": 2\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/proxy.go",
    "groupTitle": "proxy",
    "name": "GetProxylist"
  },
  {
    "type": "Post",
    "url": "/addProxy",
    "title": "添加代理服务器",
    "description": "<p>添加代理服务器</p>",
    "group": "proxy",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "addr",
            "description": "<p>IP地址</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "port",
            "description": "<p>端口</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user_name",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/proxy.go",
    "groupTitle": "proxy",
    "name": "PostAddproxy"
  },
  {
    "type": "Put",
    "url": "/saveProxy/{id}",
    "title": "修改保存代理服务器",
    "description": "<p>修改保存代理服务器</p>",
    "group": "proxy",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "addr",
            "description": "<p>IP地址</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "port",
            "description": "<p>端口</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user_name",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/proxy.go",
    "groupTitle": "proxy",
    "name": "PutSaveproxyId"
  },
  {
    "type": "Delete",
    "url": "/deleteWhite/{id}",
    "title": "删除白名单",
    "description": "<p>删除白名单</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "DeleteDeletewhiteId"
  },
  {
    "type": "Get",
    "url": "/down",
    "title": "下载导入文件",
    "description": "<p>下载导入文件</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "file_name",
            "description": "<p>文件名</p>"
          }
        ]
      }
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "GetDown"
  },
  {
    "type": "Get",
    "url": "/getImportLog",
    "title": "获取导入记录",
    "description": "<p>获取导入记录</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "per_page",
            "description": "<p>每页显示数据条数，默认15条</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>当前的所在页码，默认第1页</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"current_page\": 1,\n\t\"data\": [\n\t {\n\t\t\"id\": 17,\n\t\t\"file_name\": \"cIBJPVRSH8TLFv4foRKQ001.xlsx\",\n\t\t\"succeed_number\": 0,\n\t\t\"failure_number\": 3,\n\t\t\"create_date\": \"2017-11-08 11:04:52\"\n\t  },\n\t],\n\t\"last_page\": 1,\n\t\"per_page\": 15,\n\t\"total\": 2\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "GetGetimportlog"
  },
  {
    "type": "Get",
    "url": "/showWhite/{id}",
    "title": "修改时获取白名单信息",
    "description": "<p>修改时获取白名单信息</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"id\": 54,\n\t\"domain\": \"www.baidu002.com\",\n\t\"hall_name\": \"baidu002\",\n\t\"status\": 1,\n\t\"channel\": 1,\n\t\"ips\" : \"\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "GetShowwhiteId"
  },
  {
    "type": "Get",
    "url": "/whiteList",
    "title": "白名单列表",
    "description": "<p>白名单列表</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "search",
            "description": "<p>输入框条件</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "channel",
            "description": "<p>查询模式 1为IP，2为代理</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态，2为锁定，1为启用状态，默认为0（全部）</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "per_page",
            "description": "<p>每页显示数据条数，默认15条</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>当前的所在页码，默认第1页</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"current_page\": 1,\n\t\"data\": [\n\t  {\n\t\t\"id\": 55,\n\t\t\"domain\": \"www.alibaba.com\",\n\t\t\"hall_name\": \"马云\",\n\t\t\"status\": 1, //启用状态，2为锁定，1为启用状态，默认为1\n\t\t\"channel\": 1, //1：走IP(默认)，2：走代理\n\t\t\"ips\" : \"\",  //IP字段，用英文半角分号隔开\n\t\t\"create_date\": \"0000-00-00 00:00:00\" //创建时间\n\t  },\n\t],\n\t\"last_page\": 1,\n\t\"per_page\": 15,\n\t\"total\": 2\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "GetWhitelist"
  },
  {
    "type": "Post",
    "url": "/addWhite",
    "title": "添加白名单",
    "description": "<p>添加白名单</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "domain",
            "description": "<p>域名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "hall_name",
            "description": "<p>所属人</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ips",
            "description": "<p>IP</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "channel",
            "description": "<p>1：走IP(默认)，2：走代理</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>1：启用(默认)，2：锁定</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "PostAddwhite"
  },
  {
    "type": "Post",
    "url": "/importWhite",
    "title": "导入白名单",
    "description": "<p>导入白名单</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "File",
            "optional": false,
            "field": "uploadfile",
            "description": "<p>上传文件</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功! 总共 3 条记录！ 成功导入 0 条记录; 失败 3 条记录！\",\n  \"result\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "PostImportwhite"
  },
  {
    "type": "Put",
    "url": "/saveWhite/{id}",
    "title": "编辑保存白名单",
    "description": "<p>编辑保存白名单</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "domain",
            "description": "<p>域名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "hall_name",
            "description": "<p>所属人</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ips",
            "description": "<p>IP</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "channel",
            "description": "<p>1：走IP(默认)，2：走代理</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "PutSavewhiteId"
  },
  {
    "type": "Put",
    "url": "/saveWhiteStatus/{id}",
    "title": "修改白名单状态",
    "description": "<p>修改白名单状态</p>",
    "group": "white",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态值，2为锁定，1为启用状态</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "lock_remark",
            "description": "<p>锁定原因</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/white.go",
    "groupTitle": "white",
    "name": "PutSavewhitestatusId"
  }
] });
