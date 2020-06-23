/*
 * Copyright 2003-2006, 2009, 2017, United States Government, as represented by the Administrator of the
 * National Aeronautics and Space Administration. All rights reserved.
 *
 * The NASAWorldWind/WebWorldWind platform is licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
import KmlElements from '../KmlElements';
import KmlObject from '../KmlObject';
	/**
	 * @augments KmlObject
	 * @param options
	 * @constructor
	 * @alias KmlDelete
	 */
	var KmlDelete = function(options) {
		KmlObject.call(this, options);
	};

	KmlDelete.prototype = Object.create(KmlObject.prototype);

	Object.defineProperties(KmlDelete.prototype, {
		/**
		 * All shapes which should be deleted
		 * @memberof KmlDelete.prototype
		 * @readonly
		 * @type {KmlObject[]}
		 */
		shapes: {
			get: function(){
				return this._factory.all(this);
			}
		}
	});

	/**
	 * @inheritDoc
	 */
	KmlDelete.prototype.getTagNames = function() {
		return ['Delete'];
	};

	KmlElements.addKey(KmlDelete.prototype.getTagNames()[0], KmlDelete);

	export default KmlDelete;