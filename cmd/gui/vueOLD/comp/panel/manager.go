package panel

import "github.com/p9c/pod/cmd/gui/vue/mod"

func Manager() mod.DuoVUEcomp {
	return mod.DuoVUEcomp{
		IsApp:    true,
		Name:     "Manager",
		ID:       "panelmanager",
		Version:  "0.0.1",
		CompType: "core",
		SubType:  "manager",
		Js: `
		data () {
			return {
				items: system.data.repo,
				translations: {
					active: 'active',
					activate: 'activate',
					install: 'install',
					update: 'update',
					uninstall: 'uninstall'
				},
				delay: 3000,
				modules: [],
				states: null,
				delay: 1000,
				default: {
					name: '',
					ver: null,
					desc: '',
					state: 'uninstalled',
					image: 'https://vueschool.io/media/cb103360d9341bd311b96974a83a3a91/vuejs-components-fundamentals.png',
					url: ''
			  	},
			  	nexState: {
					active: 'inactive',
					inactive: 'active',
					outdated: 'update',
					update: 'active',
					install: 'inactive',
					uninstall: 'uninstalled',
					uninstalled: 'install',
					installed: 'uninstall'
			  	}
			}
		},
		 mounted () {
				this.normalizeParams(this)
				this.normalizeTranslations()
			},
		methods:{
			normalizeParams: 
				function normalizeParams (self) {
					this.items.map(function (obj, index) {
						obj.index = index
						obj.name = self.getParam(obj, 'name');
						obj.ver = self.getParam(obj, 'ver');
						obj.desc = self.getParam(obj, 'desc');
						obj.state = self.getParam(obj, 'state');
						obj.url = self.getParam(obj, 'url');
						obj.image = self.getParam(obj, 'image');
						return obj
					})
					this.modules = this.items
				},
			normalizeTranslations: 
				function normalizeTranslations () {
					this.states = {
						active: this.getTranslation('active'),
						inactive: this.getTranslation('activate'),
						outdated: this.getTranslation('update'),
						update: this.getTranslation('active'),
						uninstalled: this.getTranslation('install'),
						uninstall: this.getTranslation('uninstall')
					}
				},
			getParam: 
				function getParam (obj, index) {
					return typeof obj[index] !== 'undefined' ? obj[index] : this.default[index]
				},
			getTranslation: 
				function getTranslation (index) {
					let translations = this.translations ? this.translations : {}
					return typeof translations[index] !== 'undefined' ? translations[index] : index
				},
			getButtonText: 
				function getButtonText (place, state, index) {
					switch (place) {
						case 'right':
						return this.isUninstall(index) || this.isUninstalled(index) ? '' : state
						case 'left':
						return this.isUpdate(index) || this.isInstall(index) || this.isUninstall(index) ? '' : state
						default:
					}
				},
			changeState: 
				function changeState (index) {
					var actualState = this.items[index].state
					var state = this.nexState[actualState]
					this.items[index].state = state
					this.$emit('on-change', index, this)
					this.$emit('on-' + state, index, this)
				},
			uninstall: 
				function uninstall (index) {
					this.items[index].state = 'installed'
					this.changeState(index)
				},
			isInstall: 
				function isInstall (index) {
					return this.items[index].state === 'install'
				},
			isUninstall: 
				function isUninstall (index) {
					return this.items[index].state === 'uninstall'
				},
			isInactive: 
				function isInactive (index) {
					return this.items[index].state === 'inactive'
				},
			isOutdated: 
				function isOutdated (index) {
					return this.items[index].state === 'outdated'
				},
			isUpdate: 
				function isUpdate (index) {
					return this.items[index].state === 'update'
				},
			isUninstalled: 
				function isUninstalled (index) {
					return this.items[index].state === 'uninstalled'
				},
			onChange: function onChange (index, vm) {
				console.info('@onChange: ' + index)
				},
			onActive: function onActive (index, vm) {
				console.info('@onActive: ' + index)
				},
				onInactive: function onInactive (index, vm) {
			console.info('@onInactive: ' + index)
				},
			onUpdate: function onUpdate (index, vm) {
				let self = this
				console.info('@onUpdate: ' + index)
			setTimeout(function () {
					vm.changeState(index)
					self.items[index].ver = '0.0.1'
					self.items[index].name = 'DUO VUE System'
					self.items[index].desc = 'DUO VUE System.'
					self.items[index].image = 'duo.jpg'
				}, this.delay)
				},
			onInstall: function onInstall (index, vm) {
				console.info('@onInstall: ' + index)
					setTimeout(function () {
						vm.changeState(index)
					}, this.delay)
				},
			onUninstall: function onUninstall (index, vm) {
				console.info('@onUninstall: ' + index)
					setTimeout(function () {
						vm.changeState(index)
					}, this.delay)
				},
			onUninstalled: function onUninstalled (index, vm) {
				console.info('@onUninstalled:'  + index)
				}
			},
		watch:{
			items: function items () {
				this.normalizeParams(this)
			},
			translations: function translations () {
				this.normalizeTranslations()
			}
		}
		`,
		Template: `
		<div id="vue-module-manager" class="rwrap">
		<div><h1 class="vmm-cell vmm-cell-12 vmm-grid">Components manager</h1>
			<div class="vmm-cell vmm-cell-12 vmm-grid">
			<div
				:id="'vmm-module-' + (index + 1)"
				v-for="(module, index) in modules" :key="module.index"
				class="vmm-module vmm-cell"
			>
				<transition name="fade" mode="out-in">
				<div
					:key="module.image"
					:style="'background: url(' + module.image + ') center / cover;'"
					class="vmm-title"
				>
					<a :href="module.url" title="" target="_blank">
					<transition name="fade" mode="out-in">
						<div :key="module.ver" v-show="module.ver" class="vmm-ver">
						{{ module.ver }}
						</div>
					</transition>
					</a>
				</div>
				</transition>
				<div v-show="module.image !== ''" class="vmm-module-border"></div>
				<div class="vmm-content">
				<transition name="fade" mode="out-in">
					<h2 :key="module.name" class="vmm-h2">{{ module.name }}</h2>
				</transition>
				<transition name="fade" mode="out-in">
					<div class="vmm-desc" :key="module.desc">
					{{ module.desc }}
					</div>
				</transition>
				<div class="vmm-list-item">
					<div class="vmm-state-btn">
					<transition name="fade" mode="out-in">
						<div :key="module.state" class="vmm-state-buttons">
						<button
							@click="changeState(index)"
							:disabled="isUpdate(index) || isInstall(index) || isUninstall(index)"
							:class="'vmm-btn vmm-btn-left vmm-btn-' + module.state"
						>
							<span
							v-show="isUpdate(index) || isInstall(index)"
							class="vmm-loader vmm-load-blue"
							></span>
							{{ getButtonText('left', states[module.state], index) }}
						</button>
						<button
							v-show="isInactive(index) || isUninstall(index)"
							@click="uninstall(index)"
							:disabled="isUninstall(index) || isUninstalled(index)"
							:class="'vmm-btn vmm-btn--uninstall'"
						>
							<span
							v-show="isUninstall(index) || isUninstalled(index)"
							class="vmm-loader vmm-load-red
							"></span>
							<span class="vmm-text-btn">
							{{ getButtonText('right', states.uninstall, index) }}
							</span>
							<span class="vmm-delete-icon">
							<svg aria-labelledby="simpleicons-github-icon" role="img" viewBox="0 0 30 30" xmlns="http://www.w3.org/2000/svg"><title id="simpleicons-delete" lang="en">Delete icon</title><path fill="#FFFFFF" d="M6 32h20l2-22h-24zM20 4v-4h-8v4h-10v6l2-2h24l2 2v-6h-10zM18 4h-4v-2h4v2z"></path></svg>
							</span>
						</button>
						</div>
					</transition>
					</div>
				</div>
				</div>
			</div>
			</div>
		</div></div>
`,
		Css: `
		#vue-module-manager {
			font-family: 'Roboto', 'sans-serif';
		  }
		  .vmm-btn {
			border: none;
			border-radius: 2px;
			height: 36px;
			padding: 0 16px;
			text-transform: uppercase;
			letter-spacing: 0;
			float: left;
			overflow: hidden;
			will-change: box-shadow;
			transition: box-shadow 0.2s cubic-bezier(0.4, 0, 1, 1), background-color 0.2s cubic-bezier(0.4, 0, 0.2, 1), color 0.2s cubic-bezier(0.4, 0, 0.2, 1), -webkit-box-shadow 0.2s cubic-bezier(0.4, 0, 1, 1);
			outline: none;
			cursor: pointer;
			text-decoration: none;
			text-align: center;
			line-height: 36px;
			background: rgba(158, 158, 158, 0.2);
			box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
			color: #fff;
			background-color: #262b36;
			font-size: 0.84rem;
		  }
		  .vmm-btn:hover {
			opacity: 0.9;
		  }
		  .vmm-btn[disabled][disabled] {
			cursor: not-allowed;
			background-color: #f1f1f1;
		  }
		  .vmm-btn-left {
			width: 110.03px;
			min-width: 110.03px;
		  }
		  .vmm-btn-active {
			background-color: #009688;
		  }
		  .vmm-btn--uninstall {
			background-color: #d91e18;
			margin-left: 8px;
			min-width: 123.59px;
		  }
		  .vmm-btn-outdated,
		  .vmm-btn-uninstalled {
			background-color: #0073aa;
			width: initial;
		  }
		  .vmm-module {
			display: -ms-flexbox;
			display: -webkit-box;
			display: flex;
			-ms-flex-direction: column;
			-webkit-box-orient: vertical;
			-webkit-box-direction: normal;
			flex-direction: column;
			font-size: 16px;
			font-weight: 400;
			min-height: 200px;
			overflow: hidden;
			width: 330px;
			z-index: 1;
			position: relative;
			background: #fff;
			border-radius: 2px;
			-webkit-box-sizing: border-box;
			box-sizing: border-box;
			-webkit-box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
			box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
		  }
		  .vmm-module:hover {
			-webkit-box-shadow: 0 8px 10px 1px rgba(0, 0, 0, 0.14), 0 3px 14px 2px rgba(0, 0, 0, 0.12), 0 5px 5px -3px rgba(0, 0, 0, 0.2);
			box-shadow: 0 8px 10px 1px rgba(0, 0, 0, 0.14), 0 3px 14px 2px rgba(0, 0, 0, 0.12), 0 5px 5px -3px rgba(0, 0, 0, 0.2);
		  }
		  .vmm-content {
			color: rgba(0, 0, 0, 0.54);
			font-size: 1rem;
			cursor: default;
			line-height: 18px;
			overflow: hidden;
			padding: 16px;
			width: 90%;
		  }
		  .vmm-module-border {
			border-top: 1px solid rgba(0, 0, 0, 0);
			width: 100% !important;
		  }
		  .vmm-module > .vmm-title {
			color: #fff;
			height: 40px;
		  }
		  .vmm-module:hover .vmm-ver {
			display: block;
		  }
		  .vmm-load-blue {
			border: 5px solid #0073aa;
		  }
		  .vmm-load-red {
			border: 5px solid #d91e18;
		  }
		  .vmm-loader {
			display: inline-block;
			width: 14px;
			height: 14px;
			margin-top: 6px;
			border-radius: 100%;
			border-top: 5px solid transparent;
			-webkit-animation: load-animate-data-v-30a83a23 infinite linear 1s;
			animation: load-animate-data-v-30a83a23 infinite linear 1s;
		  }
		  .vmm-title {
			-ms-flex-align: center;
			-webkit-box-align: center;
			align-items: center;
			color: #262b36;
			display: block;
			display: -ms-flexbox;
			display: -webkit-box;
			display: flex;
			-ms-flex-pack: stretch;
			-webkit-box-pack: stretch;
			justify-content: stretch;
			line-height: normal;
			padding: 16px;
			-webkit-perspective-origin: 165px 56px;
			perspective-origin: 165px 56px;
			-webkit-transform-origin: 165px 56px;
			transform-origin: 165px 56px;
			-webkit-box-sizing: border-box;
			box-sizing: border-box;
			-ms-flex-positive: 1;
			-webkit-box-flex: 1;
			flex-grow: 1;
		  }
		  .vmm-h2 {
			-webkit-align-self: flex-end;
			-ms-flex-item-align: end;
			align-self: flex-end;
			color: inherit;
			display: block;
			display: -ms-flexbox;
			display: -webkit-box;
			display: flex;
			font-size: 24px;
			font-weight: 300;
			line-height: normal;
			overflow: hidden;
			-webkit-transform-origin: 149px 48px;
			transform-origin: 149px 48px;
			margin: 0;
			color: #23282d;
			margin-bottom: 16px;
		  }
		  .vmm-ver {
			display: none;
			top: 8px;
			position: absolute;
			color: #fff;
			cursor: pointer;
			right: 8px;
			border-radius: 3px;
			background: #1abc9c;
			padding: 6px;
		  }
		  .vmm-ver:hover {
			opacity: 0.9;
		  }
		  .vmm-list-item {
			font-weight: 400;
			line-height: 24px;
			letter-spacing: 0.04em;
			line-height: 1;
			display: -ms-flexbox;
			display: -webkit-box;
			display: flex;
			min-height: 48px;
			-webkit-box-sizing: border-box;
			box-sizing: border-box;
			-ms-flex-direction: row;
			-webkit-box-orient: horizontal;
			-webkit-box-direction: normal;
			flex-direction: row;
			-ms-flex-wrap: nowrap;
			flex-wrap: nowrap;
			-ms-flex-align: center;
			-webkit-box-align: center;
			align-items: center;
			padding: 16px;
			cursor: default;
			color: rgba(0, 0, 0, 0.87);
			overflow: hidden;
		  }
		  .vmm-grid {
			display: -ms-flexbox;
			display: -webkit-box;
			display: flex;
			-ms-flex-flow: row wrap;
			-webkit-box-orient: horizontal;
			-webkit-box-direction: normal;
			flex-flow: row wrap;
			margin: 0 auto;
			-ms-flex-align: stretch;
			-webkit-box-align: stretch;
			align-items: stretch;
		  }
		  .vmm-cell {
			-webkit-box-sizing: border-box;
			box-sizing: border-box;
		  }
		  .vmm-state-btn {
			margin: 5px -23px -13px -16px;
			display: -webkit-box;
			display: -webkit-inline-box;
			display: contents;
		  }
		  .vmm-delete-icon {
			display: none;
		  }
		  .vmm-delete-icon svg {
			display: inline-block;
			width: 17px;
			height: 17px;
			display: none;
		  }
		  .fade-enter-active, .fade-leave-active {
			transition: opacity .2s ease
		  }
		  .fade-enter, .fade-enter-active, .fade-leave-active {
			opacity: 0
		  }
		  @-webkit-keyframes load-animate-data-v-30a83a23 {
			0% {
			  -webkit-transform: rotate(0deg);
			  transform: rotate(0deg);
			}
			50% {
			  -webkit-transform: rotate(180deg);
			  transform: rotate(180deg);
			  opacity: 0.35;
			}
			100% {
			  -webkit-transform: rotate(360deg);
			  transform: rotate(360deg);
			}
		  }
		  @keyframes load-animate-data-v-30a83a23 {
			0% {
			  -webkit-transform: rotate(0deg);
			  transform: rotate(0deg);
			}
			50% {
			  -webkit-transform: rotate(180deg);
			  transform: rotate(180deg);
			  opacity: 0.35;
			}
			100% {
			  -webkit-transform: rotate(360deg);
			  transform: rotate(360deg);
			}
		  }
		  @media (min-width: 480px) and (max-width: 615px) {
			.vmm-text-btn {
			  display: none;
			}
			.vmm-btn--uninstall {
			  min-width: auto;
			}
			.vmm-delete-icon svg,
			.vmm-delete-icon {
			  display: block;
			}
		  }
		  @media (max-width: 479px) {
			.vmm-grid {
			  padding: 8px;
			}
			.vmm-cell {
			  margin: 8px;
			  width: calc(100% - 16px);
			}
			.vmm-cell-12 {
			  width: calc(100% - 16px);
			}
		  }
		  @media (min-width: 480px) and (max-width: 839px) {
			.vmm-grid {
			  padding: 8px;
			}
			.vmm-cell {
			  margin: 8px;
			  width: calc(50% - 16px);
			}
			.vmm-cell-12 {
			  width: calc(100% - 16px);
			}
		  }
		  @media (min-width: 840px) and (max-width: 900px) {
			.vmm-text-btn {
			  display: none;
			}
			.vmm-btn--uninstall {
			  min-width: auto;
			}
			.vmm-delete-icon svg,
			.vmm-delete-icon {
			  display: block;
			}
		  }
		  @media (max-width: 782px) {
			.vmm-grid {
			  margin-left: 9px !important;
			}
		  }
		  @media (min-width: 840px) {
			.vmm-grid {
			  padding: 8px;
			}
			.vmm-cell {
			  margin: 8px;
			  width: calc(33.3333333333% - 16px);
			}
			.vmm-cell-12 {
			  width: calc(100% - 16px);
			}
		  }
		`,
	}
}
