export namespace app {
	
	export class SerialChoice {
	    label: string;
	    value: string;
	    disabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SerialChoice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.label = source["label"];
	        this.value = source["value"];
	        this.disabled = source["disabled"];
	    }
	}
	export class SerialOptions {
	    type: string;
	    label: string;
	    key: string;
	    children: SerialChoice[];
	
	    static createFrom(source: any = {}) {
	        return new SerialOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.label = source["label"];
	        this.key = source["key"];
	        this.children = this.convertValues(source["children"], SerialChoice);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

