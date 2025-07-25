// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.898
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Slider(activeIndex int, total int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = SliderScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = SliderStyle().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"slider\"><div class=\"slider__buttons\"><button class=\"slider__button slider__button--left\"><img src=\"/public/icons/arrow-left.svg\"></button> <button class=\"slider__button slider__button--right\"><img src=\"/public/icons/arrow-right.svg\"></button></div><div class=\"slider__content-wrapper\"><div class=\"slider__content js-slider-track\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></div><div class=\"slider__indicators\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 0; i < total; i++ {
			if i == activeIndex {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div class=\"slider__dot active\" data-index=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(i)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout/slider.templ`, Line: 26, Col: 65}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\"></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<div class=\"slider__dot\" data-index=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(i)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout/slider.templ`, Line: 28, Col: 58}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func SliderScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<script>\n    document.addEventListener(\"DOMContentLoaded\", () => {\n        const slider = document.querySelector('.slider');\n        const track = slider.querySelector('.js-slider-track');\n        const btnLeft = slider.querySelector('.slider__button--left');\n        const btnRight = slider.querySelector('.slider__button--right');\n        const dots = slider.querySelectorAll('.slider__dot');\n        const totalSlides = dots.length;\n        let currentIndex = 0;\n\n        function updateSlider() {\n            track.style.transform = `translateX(-${currentIndex * 100}%)`;\n            dots.forEach((dot, i) => {\n                dot.classList.toggle('active', i === currentIndex);\n            });\n        }\n\n        btnLeft.addEventListener('click', () => {\n            currentIndex = (currentIndex - 1 + totalSlides) % totalSlides;\n            updateSlider();\n        });\n\n        btnRight.addEventListener('click', () => {\n            currentIndex = (currentIndex + 1) % totalSlides;\n            updateSlider();\n        });\n        \n        dots.forEach(dot => {\n            dot.addEventListener('click', (e) => {\n                const index = parseInt(e.target.getAttribute('data-index'), 10);\n                if (!isNaN(index)) {\n                    currentIndex = index;\n                    updateSlider();\n                }\n            });\n        });\n\n\n        updateSlider();\n    });\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func SliderStyle() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<style>\n        .slider {\n            position: relative;\n            width: 100%;\n            height: 100%;\n            border-radius: 12px;\n            overflow: hidden;\n        }\n\n        .slider__content {\n            width: 100%;\n            height: 100%;\n        }\n\n        .slider__buttons {\n            position: absolute;\n            top: 50%;\n            left: 0;\n            width: 100%;\n            display: flex;\n            justify-content: space-between;\n            transform: translateY(-50%);\n            padding: 0 16px;\n            pointer-events: none;\n            z-index: 10;\n        }\n\n        .slider__button {\n            width: 40px;\n            height: 40px;\n            border: none;\n            border-radius: 12px;\n            background: var(--color-gray);\n            display: flex;\n            align-items: center;\n            justify-content: center;\n            cursor: pointer;\n            pointer-events: all;\n        }\n\n        .slider__button img {\n            width: 24px;\n            height: 24px;\n        }\n        .slider__indicators {\n            position: absolute;\n            bottom: 24px;\n            right: 24px;\n            display: flex;\n            gap: 8px;\n            pointer-events: none;\n        }\n\n        .slider__dot {\n            width: 10px;\n            height: 10px;\n            border-radius: 50px;\n            background-color: var(--color-white-50);\n            pointer-events: auto;\n            cursor: pointer;\n        }\n\n        .slider__dot.active {\n            background-color: var(--color-white);\n            width: 24px;\n        }\n        .slider__content-wrapper {\n            width: 100%;\n            height: 100%;\n            overflow: hidden;\n        }\n\n        .slider__content {\n            display: flex;\n            transition: transform 0.2s ease;\n            height: 100%;\n        }\n\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
