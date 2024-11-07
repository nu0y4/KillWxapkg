# 感谢与介绍

首先，特别感谢 **Ackites** 师傅的无私奉献，他发布的 [KillWxapkg 项目](https://github.com/Ackites/KillWxapkg) 为我们带来了极大的便利 🎉。

另外，也非常感谢 **ekkoo** 师傅提供的宝贵思路 💡，让我能够进一步优化工具的功能。

## 工具介绍

KillWxapkg 是一个非常强大的工具，它可以：

- ⏰ **实时检测** 微信小程序的更新。
- 📦 **自动解包** 小程序包，获取其中的资源和代码。
- 🌐 **开启 web 服务**，让你能够通过浏览器轻松查看解包内容。
- 🔍 **配合 Burp Suite 插件 HAE**，你可以方便地提取和分析小程序中的敏感信息，提升安全测试的效率。

这个工具对于安全研究和分析微信小程序中的数据有非常大的帮助，特别是敏感信息提取的场景。感谢这些大佬的无私分享，让我们能够在小程序安全研究中事半功倍！🚀

# 使用
```
KillWxapkg.exe -wxp "<根据自己微信小程序地址填写>\WeChat Files\Applet"
```

## 结合油猴脚本实现一键遍历打开所有数据文件:
```
// ==UserScript==
// @name         开炮!
// @version      1.0
// @description  当然是！开炮！
// @author       面包狗
// @match        *://localhost:1549/*
// ==/UserScript==

(function() {
    'use strict';

    // 创建按钮
    const button = document.createElement('button');
    button.textContent = '开炮!';
    button.style.position = 'fixed';
    button.style.bottom = '20px';
    button.style.right = '20px';
    button.style.backgroundColor = '#007bff';
    button.style.color = '#ffffff';
    button.style.border = 'none';
    button.style.borderRadius = '5px';
    button.style.padding = '10px 20px';
    button.style.cursor = 'pointer';
    button.style.boxShadow = '0px 4px 8px rgba(0, 0, 0, 0.2)';
    button.style.fontSize = '14px';
    button.style.zIndex = '9999';

    // 按钮点击事件
    button.addEventListener('click', function() {
        // 获取所有 a 标签
        const links = document.querySelectorAll('a');
        links.forEach(link => {
            // 检查链接是否存在 href 属性并非空链接
            if (link.href && link.href !== '#') {
                // 使用 window.open 在新窗口中打开链接
                window.open(link.href, '_blank');
            }
        });
    });

    // 将按钮添加到页面中
    document.body.appendChild(button);
})();

```

#### 注：不进行任何引流，仅为了实现一些特别的功能而已，以后更不更新随缘
#### 注：此工具为了让大家更多关注原项目，删除了绝大多数原本KillWxapkg项目的功能，只新增了一个实时检测功能
