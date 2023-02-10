// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const rental = $root.rental = (() => {

    /**
     * Namespace rental.
     * @exports rental
     * @namespace
     */
    const rental = {};

    rental.v1 = (function() {

        /**
         * Namespace v1.
         * @memberof rental
         * @namespace
         */
        const v1 = {};

        v1.Loaction = (function() {

            /**
             * Properties of a Loaction.
             * @memberof rental.v1
             * @interface ILoaction
             * @property {number|null} [latitude] Loaction latitude
             * @property {number|null} [longitude] Loaction longitude
             */

            /**
             * Constructs a new Loaction.
             * @memberof rental.v1
             * @classdesc Represents a Loaction.
             * @implements ILoaction
             * @constructor
             * @param {rental.v1.ILoaction=} [properties] Properties to set
             */
            function Loaction(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Loaction latitude.
             * @member {number} latitude
             * @memberof rental.v1.Loaction
             * @instance
             */
            Loaction.prototype.latitude = 0;

            /**
             * Loaction longitude.
             * @member {number} longitude
             * @memberof rental.v1.Loaction
             * @instance
             */
            Loaction.prototype.longitude = 0;

            /**
             * Creates a new Loaction instance using the specified properties.
             * @function create
             * @memberof rental.v1.Loaction
             * @static
             * @param {rental.v1.ILoaction=} [properties] Properties to set
             * @returns {rental.v1.Loaction} Loaction instance
             */
            Loaction.create = function create(properties) {
                return new Loaction(properties);
            };

            /**
             * Encodes the specified Loaction message. Does not implicitly {@link rental.v1.Loaction.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.Loaction
             * @static
             * @param {rental.v1.ILoaction} message Loaction message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Loaction.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.latitude != null && Object.hasOwnProperty.call(message, "latitude"))
                    writer.uint32(/* id 1, wireType 1 =*/9).double(message.latitude);
                if (message.longitude != null && Object.hasOwnProperty.call(message, "longitude"))
                    writer.uint32(/* id 2, wireType 1 =*/17).double(message.longitude);
                return writer;
            };

            /**
             * Encodes the specified Loaction message, length delimited. Does not implicitly {@link rental.v1.Loaction.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.Loaction
             * @static
             * @param {rental.v1.ILoaction} message Loaction message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Loaction.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Loaction message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.Loaction
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.Loaction} Loaction
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Loaction.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.Loaction();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.latitude = reader.double();
                        break;
                    case 2:
                        message.longitude = reader.double();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Loaction message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.Loaction
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.Loaction} Loaction
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Loaction.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Loaction message.
             * @function verify
             * @memberof rental.v1.Loaction
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Loaction.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.latitude != null && message.hasOwnProperty("latitude"))
                    if (typeof message.latitude !== "number")
                        return "latitude: number expected";
                if (message.longitude != null && message.hasOwnProperty("longitude"))
                    if (typeof message.longitude !== "number")
                        return "longitude: number expected";
                return null;
            };

            /**
             * Creates a Loaction message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.Loaction
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.Loaction} Loaction
             */
            Loaction.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.Loaction)
                    return object;
                let message = new $root.rental.v1.Loaction();
                if (object.latitude != null)
                    message.latitude = Number(object.latitude);
                if (object.longitude != null)
                    message.longitude = Number(object.longitude);
                return message;
            };

            /**
             * Creates a plain object from a Loaction message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.Loaction
             * @static
             * @param {rental.v1.Loaction} message Loaction
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Loaction.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.latitude = 0;
                    object.longitude = 0;
                }
                if (message.latitude != null && message.hasOwnProperty("latitude"))
                    object.latitude = options.json && !isFinite(message.latitude) ? String(message.latitude) : message.latitude;
                if (message.longitude != null && message.hasOwnProperty("longitude"))
                    object.longitude = options.json && !isFinite(message.longitude) ? String(message.longitude) : message.longitude;
                return object;
            };

            /**
             * Converts this Loaction to JSON.
             * @function toJSON
             * @memberof rental.v1.Loaction
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Loaction.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Loaction;
        })();

        v1.LocationStatus = (function() {

            /**
             * Properties of a LocationStatus.
             * @memberof rental.v1
             * @interface ILocationStatus
             * @property {rental.v1.ILoaction|null} [Loaction] LocationStatus Loaction
             * @property {number|null} [feeCent] LocationStatus feeCent
             * @property {number|null} [kmDriven] LocationStatus kmDriven
             * @property {string|null} [poiName] LocationStatus poiName
             * @property {number|null} [timestampSec] LocationStatus timestampSec
             */

            /**
             * Constructs a new LocationStatus.
             * @memberof rental.v1
             * @classdesc Represents a LocationStatus.
             * @implements ILocationStatus
             * @constructor
             * @param {rental.v1.ILocationStatus=} [properties] Properties to set
             */
            function LocationStatus(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * LocationStatus Loaction.
             * @member {rental.v1.ILoaction|null|undefined} Loaction
             * @memberof rental.v1.LocationStatus
             * @instance
             */
            LocationStatus.prototype.Loaction = null;

            /**
             * LocationStatus feeCent.
             * @member {number} feeCent
             * @memberof rental.v1.LocationStatus
             * @instance
             */
            LocationStatus.prototype.feeCent = 0;

            /**
             * LocationStatus kmDriven.
             * @member {number} kmDriven
             * @memberof rental.v1.LocationStatus
             * @instance
             */
            LocationStatus.prototype.kmDriven = 0;

            /**
             * LocationStatus poiName.
             * @member {string} poiName
             * @memberof rental.v1.LocationStatus
             * @instance
             */
            LocationStatus.prototype.poiName = "";

            /**
             * LocationStatus timestampSec.
             * @member {number} timestampSec
             * @memberof rental.v1.LocationStatus
             * @instance
             */
            LocationStatus.prototype.timestampSec = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

            /**
             * Creates a new LocationStatus instance using the specified properties.
             * @function create
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {rental.v1.ILocationStatus=} [properties] Properties to set
             * @returns {rental.v1.LocationStatus} LocationStatus instance
             */
            LocationStatus.create = function create(properties) {
                return new LocationStatus(properties);
            };

            /**
             * Encodes the specified LocationStatus message. Does not implicitly {@link rental.v1.LocationStatus.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {rental.v1.ILocationStatus} message LocationStatus message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            LocationStatus.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.Loaction != null && Object.hasOwnProperty.call(message, "Loaction"))
                    $root.rental.v1.Loaction.encode(message.Loaction, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
                if (message.feeCent != null && Object.hasOwnProperty.call(message, "feeCent"))
                    writer.uint32(/* id 2, wireType 0 =*/16).int32(message.feeCent);
                if (message.kmDriven != null && Object.hasOwnProperty.call(message, "kmDriven"))
                    writer.uint32(/* id 3, wireType 1 =*/25).double(message.kmDriven);
                if (message.poiName != null && Object.hasOwnProperty.call(message, "poiName"))
                    writer.uint32(/* id 4, wireType 2 =*/34).string(message.poiName);
                if (message.timestampSec != null && Object.hasOwnProperty.call(message, "timestampSec"))
                    writer.uint32(/* id 5, wireType 0 =*/40).int64(message.timestampSec);
                return writer;
            };

            /**
             * Encodes the specified LocationStatus message, length delimited. Does not implicitly {@link rental.v1.LocationStatus.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {rental.v1.ILocationStatus} message LocationStatus message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            LocationStatus.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a LocationStatus message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.LocationStatus} LocationStatus
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LocationStatus.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.LocationStatus();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.Loaction = $root.rental.v1.Loaction.decode(reader, reader.uint32());
                        break;
                    case 2:
                        message.feeCent = reader.int32();
                        break;
                    case 3:
                        message.kmDriven = reader.double();
                        break;
                    case 4:
                        message.poiName = reader.string();
                        break;
                    case 5:
                        message.timestampSec = reader.int64();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a LocationStatus message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.LocationStatus} LocationStatus
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LocationStatus.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a LocationStatus message.
             * @function verify
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            LocationStatus.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.Loaction != null && message.hasOwnProperty("Loaction")) {
                    let error = $root.rental.v1.Loaction.verify(message.Loaction);
                    if (error)
                        return "Loaction." + error;
                }
                if (message.feeCent != null && message.hasOwnProperty("feeCent"))
                    if (!$util.isInteger(message.feeCent))
                        return "feeCent: integer expected";
                if (message.kmDriven != null && message.hasOwnProperty("kmDriven"))
                    if (typeof message.kmDriven !== "number")
                        return "kmDriven: number expected";
                if (message.poiName != null && message.hasOwnProperty("poiName"))
                    if (!$util.isString(message.poiName))
                        return "poiName: string expected";
                if (message.timestampSec != null && message.hasOwnProperty("timestampSec"))
                    if (!$util.isInteger(message.timestampSec) && !(message.timestampSec && $util.isInteger(message.timestampSec.low) && $util.isInteger(message.timestampSec.high)))
                        return "timestampSec: integer|Long expected";
                return null;
            };

            /**
             * Creates a LocationStatus message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.LocationStatus} LocationStatus
             */
            LocationStatus.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.LocationStatus)
                    return object;
                let message = new $root.rental.v1.LocationStatus();
                if (object.Loaction != null) {
                    if (typeof object.Loaction !== "object")
                        throw TypeError(".rental.v1.LocationStatus.Loaction: object expected");
                    message.Loaction = $root.rental.v1.Loaction.fromObject(object.Loaction);
                }
                if (object.feeCent != null)
                    message.feeCent = object.feeCent | 0;
                if (object.kmDriven != null)
                    message.kmDriven = Number(object.kmDriven);
                if (object.poiName != null)
                    message.poiName = String(object.poiName);
                if (object.timestampSec != null)
                    if ($util.Long)
                        (message.timestampSec = $util.Long.fromValue(object.timestampSec)).unsigned = false;
                    else if (typeof object.timestampSec === "string")
                        message.timestampSec = parseInt(object.timestampSec, 10);
                    else if (typeof object.timestampSec === "number")
                        message.timestampSec = object.timestampSec;
                    else if (typeof object.timestampSec === "object")
                        message.timestampSec = new $util.LongBits(object.timestampSec.low >>> 0, object.timestampSec.high >>> 0).toNumber();
                return message;
            };

            /**
             * Creates a plain object from a LocationStatus message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.LocationStatus
             * @static
             * @param {rental.v1.LocationStatus} message LocationStatus
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            LocationStatus.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.Loaction = null;
                    object.feeCent = 0;
                    object.kmDriven = 0;
                    object.poiName = "";
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, false);
                        object.timestampSec = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.timestampSec = options.longs === String ? "0" : 0;
                }
                if (message.Loaction != null && message.hasOwnProperty("Loaction"))
                    object.Loaction = $root.rental.v1.Loaction.toObject(message.Loaction, options);
                if (message.feeCent != null && message.hasOwnProperty("feeCent"))
                    object.feeCent = message.feeCent;
                if (message.kmDriven != null && message.hasOwnProperty("kmDriven"))
                    object.kmDriven = options.json && !isFinite(message.kmDriven) ? String(message.kmDriven) : message.kmDriven;
                if (message.poiName != null && message.hasOwnProperty("poiName"))
                    object.poiName = message.poiName;
                if (message.timestampSec != null && message.hasOwnProperty("timestampSec"))
                    if (typeof message.timestampSec === "number")
                        object.timestampSec = options.longs === String ? String(message.timestampSec) : message.timestampSec;
                    else
                        object.timestampSec = options.longs === String ? $util.Long.prototype.toString.call(message.timestampSec) : options.longs === Number ? new $util.LongBits(message.timestampSec.low >>> 0, message.timestampSec.high >>> 0).toNumber() : message.timestampSec;
                return object;
            };

            /**
             * Converts this LocationStatus to JSON.
             * @function toJSON
             * @memberof rental.v1.LocationStatus
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            LocationStatus.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return LocationStatus;
        })();

        /**
         * TripStatus enum.
         * @name rental.v1.TripStatus
         * @enum {number}
         * @property {number} TS_NOT_SPECIFIED=0 TS_NOT_SPECIFIED value
         * @property {number} IN_PROGRESS=1 IN_PROGRESS value
         * @property {number} FINISHED=2 FINISHED value
         */
        v1.TripStatus = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "TS_NOT_SPECIFIED"] = 0;
            values[valuesById[1] = "IN_PROGRESS"] = 1;
            values[valuesById[2] = "FINISHED"] = 2;
            return values;
        })();

        v1.TripEntity = (function() {

            /**
             * Properties of a TripEntity.
             * @memberof rental.v1
             * @interface ITripEntity
             * @property {string|null} [id] TripEntity id
             * @property {rental.v1.ITrip|null} [trip] TripEntity trip
             */

            /**
             * Constructs a new TripEntity.
             * @memberof rental.v1
             * @classdesc Represents a TripEntity.
             * @implements ITripEntity
             * @constructor
             * @param {rental.v1.ITripEntity=} [properties] Properties to set
             */
            function TripEntity(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * TripEntity id.
             * @member {string} id
             * @memberof rental.v1.TripEntity
             * @instance
             */
            TripEntity.prototype.id = "";

            /**
             * TripEntity trip.
             * @member {rental.v1.ITrip|null|undefined} trip
             * @memberof rental.v1.TripEntity
             * @instance
             */
            TripEntity.prototype.trip = null;

            /**
             * Creates a new TripEntity instance using the specified properties.
             * @function create
             * @memberof rental.v1.TripEntity
             * @static
             * @param {rental.v1.ITripEntity=} [properties] Properties to set
             * @returns {rental.v1.TripEntity} TripEntity instance
             */
            TripEntity.create = function create(properties) {
                return new TripEntity(properties);
            };

            /**
             * Encodes the specified TripEntity message. Does not implicitly {@link rental.v1.TripEntity.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.TripEntity
             * @static
             * @param {rental.v1.ITripEntity} message TripEntity message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            TripEntity.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
                if (message.trip != null && Object.hasOwnProperty.call(message, "trip"))
                    $root.rental.v1.Trip.encode(message.trip, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                return writer;
            };

            /**
             * Encodes the specified TripEntity message, length delimited. Does not implicitly {@link rental.v1.TripEntity.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.TripEntity
             * @static
             * @param {rental.v1.ITripEntity} message TripEntity message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            TripEntity.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a TripEntity message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.TripEntity
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.TripEntity} TripEntity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            TripEntity.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.TripEntity();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.id = reader.string();
                        break;
                    case 2:
                        message.trip = $root.rental.v1.Trip.decode(reader, reader.uint32());
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a TripEntity message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.TripEntity
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.TripEntity} TripEntity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            TripEntity.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a TripEntity message.
             * @function verify
             * @memberof rental.v1.TripEntity
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            TripEntity.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.id != null && message.hasOwnProperty("id"))
                    if (!$util.isString(message.id))
                        return "id: string expected";
                if (message.trip != null && message.hasOwnProperty("trip")) {
                    let error = $root.rental.v1.Trip.verify(message.trip);
                    if (error)
                        return "trip." + error;
                }
                return null;
            };

            /**
             * Creates a TripEntity message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.TripEntity
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.TripEntity} TripEntity
             */
            TripEntity.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.TripEntity)
                    return object;
                let message = new $root.rental.v1.TripEntity();
                if (object.id != null)
                    message.id = String(object.id);
                if (object.trip != null) {
                    if (typeof object.trip !== "object")
                        throw TypeError(".rental.v1.TripEntity.trip: object expected");
                    message.trip = $root.rental.v1.Trip.fromObject(object.trip);
                }
                return message;
            };

            /**
             * Creates a plain object from a TripEntity message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.TripEntity
             * @static
             * @param {rental.v1.TripEntity} message TripEntity
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            TripEntity.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.id = "";
                    object.trip = null;
                }
                if (message.id != null && message.hasOwnProperty("id"))
                    object.id = message.id;
                if (message.trip != null && message.hasOwnProperty("trip"))
                    object.trip = $root.rental.v1.Trip.toObject(message.trip, options);
                return object;
            };

            /**
             * Converts this TripEntity to JSON.
             * @function toJSON
             * @memberof rental.v1.TripEntity
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            TripEntity.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return TripEntity;
        })();

        v1.Trip = (function() {

            /**
             * Properties of a Trip.
             * @memberof rental.v1
             * @interface ITrip
             * @property {string|null} [accountID] Trip accountID
             * @property {string|null} [carID] Trip carID
             * @property {rental.v1.ILocationStatus|null} [start] Trip start
             * @property {rental.v1.ILocationStatus|null} [current] Trip current
             * @property {rental.v1.ILocationStatus|null} [end] Trip end
             * @property {rental.v1.TripStatus|null} [status] Trip status
             * @property {string|null} [identityId] Trip identityId
             */

            /**
             * Constructs a new Trip.
             * @memberof rental.v1
             * @classdesc Represents a Trip.
             * @implements ITrip
             * @constructor
             * @param {rental.v1.ITrip=} [properties] Properties to set
             */
            function Trip(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Trip accountID.
             * @member {string} accountID
             * @memberof rental.v1.Trip
             * @instance
             */
            Trip.prototype.accountID = "";

            /**
             * Trip carID.
             * @member {string} carID
             * @memberof rental.v1.Trip
             * @instance
             */
            Trip.prototype.carID = "";

            /**
             * Trip start.
             * @member {rental.v1.ILocationStatus|null|undefined} start
             * @memberof rental.v1.Trip
             * @instance
             */
            Trip.prototype.start = null;

            /**
             * Trip current.
             * @member {rental.v1.ILocationStatus|null|undefined} current
             * @memberof rental.v1.Trip
             * @instance
             */
            Trip.prototype.current = null;

            /**
             * Trip end.
             * @member {rental.v1.ILocationStatus|null|undefined} end
             * @memberof rental.v1.Trip
             * @instance
             */
            Trip.prototype.end = null;

            /**
             * Trip status.
             * @member {rental.v1.TripStatus} status
             * @memberof rental.v1.Trip
             * @instance
             */
            Trip.prototype.status = 0;

            /**
             * Trip identityId.
             * @member {string} identityId
             * @memberof rental.v1.Trip
             * @instance
             */
            Trip.prototype.identityId = "";

            /**
             * Creates a new Trip instance using the specified properties.
             * @function create
             * @memberof rental.v1.Trip
             * @static
             * @param {rental.v1.ITrip=} [properties] Properties to set
             * @returns {rental.v1.Trip} Trip instance
             */
            Trip.create = function create(properties) {
                return new Trip(properties);
            };

            /**
             * Encodes the specified Trip message. Does not implicitly {@link rental.v1.Trip.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.Trip
             * @static
             * @param {rental.v1.ITrip} message Trip message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Trip.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.accountID != null && Object.hasOwnProperty.call(message, "accountID"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.accountID);
                if (message.carID != null && Object.hasOwnProperty.call(message, "carID"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.carID);
                if (message.start != null && Object.hasOwnProperty.call(message, "start"))
                    $root.rental.v1.LocationStatus.encode(message.start, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
                if (message.current != null && Object.hasOwnProperty.call(message, "current"))
                    $root.rental.v1.LocationStatus.encode(message.current, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
                if (message.end != null && Object.hasOwnProperty.call(message, "end"))
                    $root.rental.v1.LocationStatus.encode(message.end, writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
                if (message.status != null && Object.hasOwnProperty.call(message, "status"))
                    writer.uint32(/* id 6, wireType 0 =*/48).int32(message.status);
                if (message.identityId != null && Object.hasOwnProperty.call(message, "identityId"))
                    writer.uint32(/* id 7, wireType 2 =*/58).string(message.identityId);
                return writer;
            };

            /**
             * Encodes the specified Trip message, length delimited. Does not implicitly {@link rental.v1.Trip.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.Trip
             * @static
             * @param {rental.v1.ITrip} message Trip message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Trip.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Trip message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.Trip
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.Trip} Trip
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Trip.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.Trip();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.accountID = reader.string();
                        break;
                    case 2:
                        message.carID = reader.string();
                        break;
                    case 3:
                        message.start = $root.rental.v1.LocationStatus.decode(reader, reader.uint32());
                        break;
                    case 4:
                        message.current = $root.rental.v1.LocationStatus.decode(reader, reader.uint32());
                        break;
                    case 5:
                        message.end = $root.rental.v1.LocationStatus.decode(reader, reader.uint32());
                        break;
                    case 6:
                        message.status = reader.int32();
                        break;
                    case 7:
                        message.identityId = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Trip message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.Trip
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.Trip} Trip
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Trip.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Trip message.
             * @function verify
             * @memberof rental.v1.Trip
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Trip.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.accountID != null && message.hasOwnProperty("accountID"))
                    if (!$util.isString(message.accountID))
                        return "accountID: string expected";
                if (message.carID != null && message.hasOwnProperty("carID"))
                    if (!$util.isString(message.carID))
                        return "carID: string expected";
                if (message.start != null && message.hasOwnProperty("start")) {
                    let error = $root.rental.v1.LocationStatus.verify(message.start);
                    if (error)
                        return "start." + error;
                }
                if (message.current != null && message.hasOwnProperty("current")) {
                    let error = $root.rental.v1.LocationStatus.verify(message.current);
                    if (error)
                        return "current." + error;
                }
                if (message.end != null && message.hasOwnProperty("end")) {
                    let error = $root.rental.v1.LocationStatus.verify(message.end);
                    if (error)
                        return "end." + error;
                }
                if (message.status != null && message.hasOwnProperty("status"))
                    switch (message.status) {
                    default:
                        return "status: enum value expected";
                    case 0:
                    case 1:
                    case 2:
                        break;
                    }
                if (message.identityId != null && message.hasOwnProperty("identityId"))
                    if (!$util.isString(message.identityId))
                        return "identityId: string expected";
                return null;
            };

            /**
             * Creates a Trip message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.Trip
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.Trip} Trip
             */
            Trip.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.Trip)
                    return object;
                let message = new $root.rental.v1.Trip();
                if (object.accountID != null)
                    message.accountID = String(object.accountID);
                if (object.carID != null)
                    message.carID = String(object.carID);
                if (object.start != null) {
                    if (typeof object.start !== "object")
                        throw TypeError(".rental.v1.Trip.start: object expected");
                    message.start = $root.rental.v1.LocationStatus.fromObject(object.start);
                }
                if (object.current != null) {
                    if (typeof object.current !== "object")
                        throw TypeError(".rental.v1.Trip.current: object expected");
                    message.current = $root.rental.v1.LocationStatus.fromObject(object.current);
                }
                if (object.end != null) {
                    if (typeof object.end !== "object")
                        throw TypeError(".rental.v1.Trip.end: object expected");
                    message.end = $root.rental.v1.LocationStatus.fromObject(object.end);
                }
                switch (object.status) {
                case "TS_NOT_SPECIFIED":
                case 0:
                    message.status = 0;
                    break;
                case "IN_PROGRESS":
                case 1:
                    message.status = 1;
                    break;
                case "FINISHED":
                case 2:
                    message.status = 2;
                    break;
                }
                if (object.identityId != null)
                    message.identityId = String(object.identityId);
                return message;
            };

            /**
             * Creates a plain object from a Trip message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.Trip
             * @static
             * @param {rental.v1.Trip} message Trip
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Trip.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.accountID = "";
                    object.carID = "";
                    object.start = null;
                    object.current = null;
                    object.end = null;
                    object.status = options.enums === String ? "TS_NOT_SPECIFIED" : 0;
                    object.identityId = "";
                }
                if (message.accountID != null && message.hasOwnProperty("accountID"))
                    object.accountID = message.accountID;
                if (message.carID != null && message.hasOwnProperty("carID"))
                    object.carID = message.carID;
                if (message.start != null && message.hasOwnProperty("start"))
                    object.start = $root.rental.v1.LocationStatus.toObject(message.start, options);
                if (message.current != null && message.hasOwnProperty("current"))
                    object.current = $root.rental.v1.LocationStatus.toObject(message.current, options);
                if (message.end != null && message.hasOwnProperty("end"))
                    object.end = $root.rental.v1.LocationStatus.toObject(message.end, options);
                if (message.status != null && message.hasOwnProperty("status"))
                    object.status = options.enums === String ? $root.rental.v1.TripStatus[message.status] : message.status;
                if (message.identityId != null && message.hasOwnProperty("identityId"))
                    object.identityId = message.identityId;
                return object;
            };

            /**
             * Converts this Trip to JSON.
             * @function toJSON
             * @memberof rental.v1.Trip
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Trip.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Trip;
        })();

        v1.CreateTripRequest = (function() {

            /**
             * Properties of a CreateTripRequest.
             * @memberof rental.v1
             * @interface ICreateTripRequest
             * @property {rental.v1.ILoaction|null} [start] CreateTripRequest start
             * @property {string|null} [carId] CreateTripRequest carId
             */

            /**
             * Constructs a new CreateTripRequest.
             * @memberof rental.v1
             * @classdesc Represents a CreateTripRequest.
             * @implements ICreateTripRequest
             * @constructor
             * @param {rental.v1.ICreateTripRequest=} [properties] Properties to set
             */
            function CreateTripRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * CreateTripRequest start.
             * @member {rental.v1.ILoaction|null|undefined} start
             * @memberof rental.v1.CreateTripRequest
             * @instance
             */
            CreateTripRequest.prototype.start = null;

            /**
             * CreateTripRequest carId.
             * @member {string} carId
             * @memberof rental.v1.CreateTripRequest
             * @instance
             */
            CreateTripRequest.prototype.carId = "";

            /**
             * Creates a new CreateTripRequest instance using the specified properties.
             * @function create
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.ICreateTripRequest=} [properties] Properties to set
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest instance
             */
            CreateTripRequest.create = function create(properties) {
                return new CreateTripRequest(properties);
            };

            /**
             * Encodes the specified CreateTripRequest message. Does not implicitly {@link rental.v1.CreateTripRequest.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.ICreateTripRequest} message CreateTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CreateTripRequest.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.start != null && Object.hasOwnProperty.call(message, "start"))
                    $root.rental.v1.Loaction.encode(message.start, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
                if (message.carId != null && Object.hasOwnProperty.call(message, "carId"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.carId);
                return writer;
            };

            /**
             * Encodes the specified CreateTripRequest message, length delimited. Does not implicitly {@link rental.v1.CreateTripRequest.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.ICreateTripRequest} message CreateTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CreateTripRequest.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a CreateTripRequest message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CreateTripRequest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.CreateTripRequest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.start = $root.rental.v1.Loaction.decode(reader, reader.uint32());
                        break;
                    case 2:
                        message.carId = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a CreateTripRequest message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CreateTripRequest.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a CreateTripRequest message.
             * @function verify
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            CreateTripRequest.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.start != null && message.hasOwnProperty("start")) {
                    let error = $root.rental.v1.Loaction.verify(message.start);
                    if (error)
                        return "start." + error;
                }
                if (message.carId != null && message.hasOwnProperty("carId"))
                    if (!$util.isString(message.carId))
                        return "carId: string expected";
                return null;
            };

            /**
             * Creates a CreateTripRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest
             */
            CreateTripRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.CreateTripRequest)
                    return object;
                let message = new $root.rental.v1.CreateTripRequest();
                if (object.start != null) {
                    if (typeof object.start !== "object")
                        throw TypeError(".rental.v1.CreateTripRequest.start: object expected");
                    message.start = $root.rental.v1.Loaction.fromObject(object.start);
                }
                if (object.carId != null)
                    message.carId = String(object.carId);
                return message;
            };

            /**
             * Creates a plain object from a CreateTripRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.CreateTripRequest} message CreateTripRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            CreateTripRequest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.start = null;
                    object.carId = "";
                }
                if (message.start != null && message.hasOwnProperty("start"))
                    object.start = $root.rental.v1.Loaction.toObject(message.start, options);
                if (message.carId != null && message.hasOwnProperty("carId"))
                    object.carId = message.carId;
                return object;
            };

            /**
             * Converts this CreateTripRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.CreateTripRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            CreateTripRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return CreateTripRequest;
        })();

        v1.GetTripRequest = (function() {

            /**
             * Properties of a GetTripRequest.
             * @memberof rental.v1
             * @interface IGetTripRequest
             * @property {string|null} [id] GetTripRequest id
             */

            /**
             * Constructs a new GetTripRequest.
             * @memberof rental.v1
             * @classdesc Represents a GetTripRequest.
             * @implements IGetTripRequest
             * @constructor
             * @param {rental.v1.IGetTripRequest=} [properties] Properties to set
             */
            function GetTripRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * GetTripRequest id.
             * @member {string} id
             * @memberof rental.v1.GetTripRequest
             * @instance
             */
            GetTripRequest.prototype.id = "";

            /**
             * Creates a new GetTripRequest instance using the specified properties.
             * @function create
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {rental.v1.IGetTripRequest=} [properties] Properties to set
             * @returns {rental.v1.GetTripRequest} GetTripRequest instance
             */
            GetTripRequest.create = function create(properties) {
                return new GetTripRequest(properties);
            };

            /**
             * Encodes the specified GetTripRequest message. Does not implicitly {@link rental.v1.GetTripRequest.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {rental.v1.IGetTripRequest} message GetTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetTripRequest.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
                return writer;
            };

            /**
             * Encodes the specified GetTripRequest message, length delimited. Does not implicitly {@link rental.v1.GetTripRequest.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {rental.v1.IGetTripRequest} message GetTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetTripRequest.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a GetTripRequest message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.GetTripRequest} GetTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetTripRequest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.GetTripRequest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.id = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a GetTripRequest message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.GetTripRequest} GetTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetTripRequest.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a GetTripRequest message.
             * @function verify
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            GetTripRequest.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.id != null && message.hasOwnProperty("id"))
                    if (!$util.isString(message.id))
                        return "id: string expected";
                return null;
            };

            /**
             * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetTripRequest} GetTripRequest
             */
            GetTripRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetTripRequest)
                    return object;
                let message = new $root.rental.v1.GetTripRequest();
                if (object.id != null)
                    message.id = String(object.id);
                return message;
            };

            /**
             * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetTripRequest
             * @static
             * @param {rental.v1.GetTripRequest} message GetTripRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetTripRequest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults)
                    object.id = "";
                if (message.id != null && message.hasOwnProperty("id"))
                    object.id = message.id;
                return object;
            };

            /**
             * Converts this GetTripRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.GetTripRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetTripRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return GetTripRequest;
        })();

        v1.GetTripsRequest = (function() {

            /**
             * Properties of a GetTripsRequest.
             * @memberof rental.v1
             * @interface IGetTripsRequest
             * @property {rental.v1.TripStatus|null} [status] GetTripsRequest status
             */

            /**
             * Constructs a new GetTripsRequest.
             * @memberof rental.v1
             * @classdesc Represents a GetTripsRequest.
             * @implements IGetTripsRequest
             * @constructor
             * @param {rental.v1.IGetTripsRequest=} [properties] Properties to set
             */
            function GetTripsRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * GetTripsRequest status.
             * @member {rental.v1.TripStatus} status
             * @memberof rental.v1.GetTripsRequest
             * @instance
             */
            GetTripsRequest.prototype.status = 0;

            /**
             * Creates a new GetTripsRequest instance using the specified properties.
             * @function create
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {rental.v1.IGetTripsRequest=} [properties] Properties to set
             * @returns {rental.v1.GetTripsRequest} GetTripsRequest instance
             */
            GetTripsRequest.create = function create(properties) {
                return new GetTripsRequest(properties);
            };

            /**
             * Encodes the specified GetTripsRequest message. Does not implicitly {@link rental.v1.GetTripsRequest.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {rental.v1.IGetTripsRequest} message GetTripsRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetTripsRequest.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.status != null && Object.hasOwnProperty.call(message, "status"))
                    writer.uint32(/* id 1, wireType 0 =*/8).int32(message.status);
                return writer;
            };

            /**
             * Encodes the specified GetTripsRequest message, length delimited. Does not implicitly {@link rental.v1.GetTripsRequest.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {rental.v1.IGetTripsRequest} message GetTripsRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetTripsRequest.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a GetTripsRequest message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.GetTripsRequest} GetTripsRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetTripsRequest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.GetTripsRequest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.status = reader.int32();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a GetTripsRequest message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.GetTripsRequest} GetTripsRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetTripsRequest.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a GetTripsRequest message.
             * @function verify
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            GetTripsRequest.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.status != null && message.hasOwnProperty("status"))
                    switch (message.status) {
                    default:
                        return "status: enum value expected";
                    case 0:
                    case 1:
                    case 2:
                        break;
                    }
                return null;
            };

            /**
             * Creates a GetTripsRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetTripsRequest} GetTripsRequest
             */
            GetTripsRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetTripsRequest)
                    return object;
                let message = new $root.rental.v1.GetTripsRequest();
                switch (object.status) {
                case "TS_NOT_SPECIFIED":
                case 0:
                    message.status = 0;
                    break;
                case "IN_PROGRESS":
                case 1:
                    message.status = 1;
                    break;
                case "FINISHED":
                case 2:
                    message.status = 2;
                    break;
                }
                return message;
            };

            /**
             * Creates a plain object from a GetTripsRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetTripsRequest
             * @static
             * @param {rental.v1.GetTripsRequest} message GetTripsRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetTripsRequest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults)
                    object.status = options.enums === String ? "TS_NOT_SPECIFIED" : 0;
                if (message.status != null && message.hasOwnProperty("status"))
                    object.status = options.enums === String ? $root.rental.v1.TripStatus[message.status] : message.status;
                return object;
            };

            /**
             * Converts this GetTripsRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.GetTripsRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetTripsRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return GetTripsRequest;
        })();

        v1.GetTripsResponse = (function() {

            /**
             * Properties of a GetTripsResponse.
             * @memberof rental.v1
             * @interface IGetTripsResponse
             * @property {Array.<rental.v1.ITripEntity>|null} [trips] GetTripsResponse trips
             */

            /**
             * Constructs a new GetTripsResponse.
             * @memberof rental.v1
             * @classdesc Represents a GetTripsResponse.
             * @implements IGetTripsResponse
             * @constructor
             * @param {rental.v1.IGetTripsResponse=} [properties] Properties to set
             */
            function GetTripsResponse(properties) {
                this.trips = [];
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * GetTripsResponse trips.
             * @member {Array.<rental.v1.ITripEntity>} trips
             * @memberof rental.v1.GetTripsResponse
             * @instance
             */
            GetTripsResponse.prototype.trips = $util.emptyArray;

            /**
             * Creates a new GetTripsResponse instance using the specified properties.
             * @function create
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {rental.v1.IGetTripsResponse=} [properties] Properties to set
             * @returns {rental.v1.GetTripsResponse} GetTripsResponse instance
             */
            GetTripsResponse.create = function create(properties) {
                return new GetTripsResponse(properties);
            };

            /**
             * Encodes the specified GetTripsResponse message. Does not implicitly {@link rental.v1.GetTripsResponse.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {rental.v1.IGetTripsResponse} message GetTripsResponse message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetTripsResponse.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.trips != null && message.trips.length)
                    for (let i = 0; i < message.trips.length; ++i)
                        $root.rental.v1.TripEntity.encode(message.trips[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
                return writer;
            };

            /**
             * Encodes the specified GetTripsResponse message, length delimited. Does not implicitly {@link rental.v1.GetTripsResponse.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {rental.v1.IGetTripsResponse} message GetTripsResponse message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetTripsResponse.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a GetTripsResponse message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.GetTripsResponse} GetTripsResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetTripsResponse.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.GetTripsResponse();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        if (!(message.trips && message.trips.length))
                            message.trips = [];
                        message.trips.push($root.rental.v1.TripEntity.decode(reader, reader.uint32()));
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a GetTripsResponse message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.GetTripsResponse} GetTripsResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetTripsResponse.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a GetTripsResponse message.
             * @function verify
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            GetTripsResponse.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.trips != null && message.hasOwnProperty("trips")) {
                    if (!Array.isArray(message.trips))
                        return "trips: array expected";
                    for (let i = 0; i < message.trips.length; ++i) {
                        let error = $root.rental.v1.TripEntity.verify(message.trips[i]);
                        if (error)
                            return "trips." + error;
                    }
                }
                return null;
            };

            /**
             * Creates a GetTripsResponse message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetTripsResponse} GetTripsResponse
             */
            GetTripsResponse.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetTripsResponse)
                    return object;
                let message = new $root.rental.v1.GetTripsResponse();
                if (object.trips) {
                    if (!Array.isArray(object.trips))
                        throw TypeError(".rental.v1.GetTripsResponse.trips: array expected");
                    message.trips = [];
                    for (let i = 0; i < object.trips.length; ++i) {
                        if (typeof object.trips[i] !== "object")
                            throw TypeError(".rental.v1.GetTripsResponse.trips: object expected");
                        message.trips[i] = $root.rental.v1.TripEntity.fromObject(object.trips[i]);
                    }
                }
                return message;
            };

            /**
             * Creates a plain object from a GetTripsResponse message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetTripsResponse
             * @static
             * @param {rental.v1.GetTripsResponse} message GetTripsResponse
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetTripsResponse.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.arrays || options.defaults)
                    object.trips = [];
                if (message.trips && message.trips.length) {
                    object.trips = [];
                    for (let j = 0; j < message.trips.length; ++j)
                        object.trips[j] = $root.rental.v1.TripEntity.toObject(message.trips[j], options);
                }
                return object;
            };

            /**
             * Converts this GetTripsResponse to JSON.
             * @function toJSON
             * @memberof rental.v1.GetTripsResponse
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetTripsResponse.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return GetTripsResponse;
        })();

        v1.UpdateTripRequest = (function() {

            /**
             * Properties of an UpdateTripRequest.
             * @memberof rental.v1
             * @interface IUpdateTripRequest
             * @property {string|null} [id] UpdateTripRequest id
             * @property {rental.v1.ILoaction|null} [current] UpdateTripRequest current
             * @property {boolean|null} [endTrip] UpdateTripRequest endTrip
             */

            /**
             * Constructs a new UpdateTripRequest.
             * @memberof rental.v1
             * @classdesc Represents an UpdateTripRequest.
             * @implements IUpdateTripRequest
             * @constructor
             * @param {rental.v1.IUpdateTripRequest=} [properties] Properties to set
             */
            function UpdateTripRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * UpdateTripRequest id.
             * @member {string} id
             * @memberof rental.v1.UpdateTripRequest
             * @instance
             */
            UpdateTripRequest.prototype.id = "";

            /**
             * UpdateTripRequest current.
             * @member {rental.v1.ILoaction|null|undefined} current
             * @memberof rental.v1.UpdateTripRequest
             * @instance
             */
            UpdateTripRequest.prototype.current = null;

            /**
             * UpdateTripRequest endTrip.
             * @member {boolean} endTrip
             * @memberof rental.v1.UpdateTripRequest
             * @instance
             */
            UpdateTripRequest.prototype.endTrip = false;

            /**
             * Creates a new UpdateTripRequest instance using the specified properties.
             * @function create
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {rental.v1.IUpdateTripRequest=} [properties] Properties to set
             * @returns {rental.v1.UpdateTripRequest} UpdateTripRequest instance
             */
            UpdateTripRequest.create = function create(properties) {
                return new UpdateTripRequest(properties);
            };

            /**
             * Encodes the specified UpdateTripRequest message. Does not implicitly {@link rental.v1.UpdateTripRequest.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {rental.v1.IUpdateTripRequest} message UpdateTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            UpdateTripRequest.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
                if (message.current != null && Object.hasOwnProperty.call(message, "current"))
                    $root.rental.v1.Loaction.encode(message.current, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                if (message.endTrip != null && Object.hasOwnProperty.call(message, "endTrip"))
                    writer.uint32(/* id 3, wireType 0 =*/24).bool(message.endTrip);
                return writer;
            };

            /**
             * Encodes the specified UpdateTripRequest message, length delimited. Does not implicitly {@link rental.v1.UpdateTripRequest.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {rental.v1.IUpdateTripRequest} message UpdateTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            UpdateTripRequest.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes an UpdateTripRequest message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.UpdateTripRequest} UpdateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            UpdateTripRequest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.UpdateTripRequest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.id = reader.string();
                        break;
                    case 2:
                        message.current = $root.rental.v1.Loaction.decode(reader, reader.uint32());
                        break;
                    case 3:
                        message.endTrip = reader.bool();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes an UpdateTripRequest message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.UpdateTripRequest} UpdateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            UpdateTripRequest.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies an UpdateTripRequest message.
             * @function verify
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            UpdateTripRequest.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.id != null && message.hasOwnProperty("id"))
                    if (!$util.isString(message.id))
                        return "id: string expected";
                if (message.current != null && message.hasOwnProperty("current")) {
                    let error = $root.rental.v1.Loaction.verify(message.current);
                    if (error)
                        return "current." + error;
                }
                if (message.endTrip != null && message.hasOwnProperty("endTrip"))
                    if (typeof message.endTrip !== "boolean")
                        return "endTrip: boolean expected";
                return null;
            };

            /**
             * Creates an UpdateTripRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.UpdateTripRequest} UpdateTripRequest
             */
            UpdateTripRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.UpdateTripRequest)
                    return object;
                let message = new $root.rental.v1.UpdateTripRequest();
                if (object.id != null)
                    message.id = String(object.id);
                if (object.current != null) {
                    if (typeof object.current !== "object")
                        throw TypeError(".rental.v1.UpdateTripRequest.current: object expected");
                    message.current = $root.rental.v1.Loaction.fromObject(object.current);
                }
                if (object.endTrip != null)
                    message.endTrip = Boolean(object.endTrip);
                return message;
            };

            /**
             * Creates a plain object from an UpdateTripRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.UpdateTripRequest
             * @static
             * @param {rental.v1.UpdateTripRequest} message UpdateTripRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            UpdateTripRequest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.id = "";
                    object.current = null;
                    object.endTrip = false;
                }
                if (message.id != null && message.hasOwnProperty("id"))
                    object.id = message.id;
                if (message.current != null && message.hasOwnProperty("current"))
                    object.current = $root.rental.v1.Loaction.toObject(message.current, options);
                if (message.endTrip != null && message.hasOwnProperty("endTrip"))
                    object.endTrip = message.endTrip;
                return object;
            };

            /**
             * Converts this UpdateTripRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.UpdateTripRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            UpdateTripRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return UpdateTripRequest;
        })();

        v1.TripService = (function() {

            /**
             * Constructs a new TripService service.
             * @memberof rental.v1
             * @classdesc Represents a TripService
             * @extends $protobuf.rpc.Service
             * @constructor
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             */
            function TripService(rpcImpl, requestDelimited, responseDelimited) {
                $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
            }

            (TripService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = TripService;

            /**
             * Creates new TripService service using the specified rpc implementation.
             * @function create
             * @memberof rental.v1.TripService
             * @static
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             * @returns {TripService} RPC service. Useful where requests and/or responses are streamed.
             */
            TripService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
                return new this(rpcImpl, requestDelimited, responseDelimited);
            };

            /**
             * Callback as used by {@link rental.v1.TripService#createTrip}.
             * @memberof rental.v1.TripService
             * @typedef CreateTripCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.TripEntity} [response] TripEntity
             */

            /**
             * Calls CreateTrip.
             * @function createTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.ICreateTripRequest} request CreateTripRequest message or plain object
             * @param {rental.v1.TripService.CreateTripCallback} callback Node-style callback called with the error, if any, and TripEntity
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(TripService.prototype.createTrip = function createTrip(request, callback) {
                return this.rpcCall(createTrip, $root.rental.v1.CreateTripRequest, $root.rental.v1.TripEntity, request, callback);
            }, "name", { value: "CreateTrip" });

            /**
             * Calls CreateTrip.
             * @function createTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.ICreateTripRequest} request CreateTripRequest message or plain object
             * @returns {Promise<rental.v1.TripEntity>} Promise
             * @variation 2
             */

            /**
             * Callback as used by {@link rental.v1.TripService#getTrip}.
             * @memberof rental.v1.TripService
             * @typedef GetTripCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.Trip} [response] Trip
             */

            /**
             * Calls GetTrip.
             * @function getTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.IGetTripRequest} request GetTripRequest message or plain object
             * @param {rental.v1.TripService.GetTripCallback} callback Node-style callback called with the error, if any, and Trip
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(TripService.prototype.getTrip = function getTrip(request, callback) {
                return this.rpcCall(getTrip, $root.rental.v1.GetTripRequest, $root.rental.v1.Trip, request, callback);
            }, "name", { value: "GetTrip" });

            /**
             * Calls GetTrip.
             * @function getTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.IGetTripRequest} request GetTripRequest message or plain object
             * @returns {Promise<rental.v1.Trip>} Promise
             * @variation 2
             */

            /**
             * Callback as used by {@link rental.v1.TripService#getTrips}.
             * @memberof rental.v1.TripService
             * @typedef GetTripsCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.GetTripsResponse} [response] GetTripsResponse
             */

            /**
             * Calls GetTrips.
             * @function getTrips
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.IGetTripsRequest} request GetTripsRequest message or plain object
             * @param {rental.v1.TripService.GetTripsCallback} callback Node-style callback called with the error, if any, and GetTripsResponse
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(TripService.prototype.getTrips = function getTrips(request, callback) {
                return this.rpcCall(getTrips, $root.rental.v1.GetTripsRequest, $root.rental.v1.GetTripsResponse, request, callback);
            }, "name", { value: "GetTrips" });

            /**
             * Calls GetTrips.
             * @function getTrips
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.IGetTripsRequest} request GetTripsRequest message or plain object
             * @returns {Promise<rental.v1.GetTripsResponse>} Promise
             * @variation 2
             */

            /**
             * Callback as used by {@link rental.v1.TripService#updateTrip}.
             * @memberof rental.v1.TripService
             * @typedef UpdateTripCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.Trip} [response] Trip
             */

            /**
             * Calls UpdateTrip.
             * @function updateTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.IUpdateTripRequest} request UpdateTripRequest message or plain object
             * @param {rental.v1.TripService.UpdateTripCallback} callback Node-style callback called with the error, if any, and Trip
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(TripService.prototype.updateTrip = function updateTrip(request, callback) {
                return this.rpcCall(updateTrip, $root.rental.v1.UpdateTripRequest, $root.rental.v1.Trip, request, callback);
            }, "name", { value: "UpdateTrip" });

            /**
             * Calls UpdateTrip.
             * @function updateTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.IUpdateTripRequest} request UpdateTripRequest message or plain object
             * @returns {Promise<rental.v1.Trip>} Promise
             * @variation 2
             */

            return TripService;
        })();

        /**
         * Gender enum.
         * @name rental.v1.Gender
         * @enum {number}
         * @property {number} G_NOT_SPECIFIED=0 G_NOT_SPECIFIED value
         * @property {number} MALE=1 MALE value
         * @property {number} FEMAEL=2 FEMAEL value
         */
        v1.Gender = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "G_NOT_SPECIFIED"] = 0;
            values[valuesById[1] = "MALE"] = 1;
            values[valuesById[2] = "FEMAEL"] = 2;
            return values;
        })();

        /**
         * IdentityStatus enum.
         * @name rental.v1.IdentityStatus
         * @enum {number}
         * @property {number} UNSUBMITTED=0 UNSUBMITTED value
         * @property {number} PENDING=1 PENDING value
         * @property {number} VERIFIED=2 VERIFIED value
         */
        v1.IdentityStatus = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "UNSUBMITTED"] = 0;
            values[valuesById[1] = "PENDING"] = 1;
            values[valuesById[2] = "VERIFIED"] = 2;
            return values;
        })();

        v1.Profile = (function() {

            /**
             * Properties of a Profile.
             * @memberof rental.v1
             * @interface IProfile
             * @property {rental.v1.IIdentity|null} [identity] Profile identity
             * @property {rental.v1.IdentityStatus|null} [identityStatus] Profile identityStatus
             */

            /**
             * Constructs a new Profile.
             * @memberof rental.v1
             * @classdesc Represents a Profile.
             * @implements IProfile
             * @constructor
             * @param {rental.v1.IProfile=} [properties] Properties to set
             */
            function Profile(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Profile identity.
             * @member {rental.v1.IIdentity|null|undefined} identity
             * @memberof rental.v1.Profile
             * @instance
             */
            Profile.prototype.identity = null;

            /**
             * Profile identityStatus.
             * @member {rental.v1.IdentityStatus} identityStatus
             * @memberof rental.v1.Profile
             * @instance
             */
            Profile.prototype.identityStatus = 0;

            /**
             * Creates a new Profile instance using the specified properties.
             * @function create
             * @memberof rental.v1.Profile
             * @static
             * @param {rental.v1.IProfile=} [properties] Properties to set
             * @returns {rental.v1.Profile} Profile instance
             */
            Profile.create = function create(properties) {
                return new Profile(properties);
            };

            /**
             * Encodes the specified Profile message. Does not implicitly {@link rental.v1.Profile.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.Profile
             * @static
             * @param {rental.v1.IProfile} message Profile message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Profile.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.identity != null && Object.hasOwnProperty.call(message, "identity"))
                    $root.rental.v1.Identity.encode(message.identity, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
                if (message.identityStatus != null && Object.hasOwnProperty.call(message, "identityStatus"))
                    writer.uint32(/* id 2, wireType 0 =*/16).int32(message.identityStatus);
                return writer;
            };

            /**
             * Encodes the specified Profile message, length delimited. Does not implicitly {@link rental.v1.Profile.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.Profile
             * @static
             * @param {rental.v1.IProfile} message Profile message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Profile.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Profile message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.Profile
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.Profile} Profile
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Profile.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.Profile();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.identity = $root.rental.v1.Identity.decode(reader, reader.uint32());
                        break;
                    case 2:
                        message.identityStatus = reader.int32();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Profile message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.Profile
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.Profile} Profile
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Profile.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Profile message.
             * @function verify
             * @memberof rental.v1.Profile
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Profile.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.identity != null && message.hasOwnProperty("identity")) {
                    let error = $root.rental.v1.Identity.verify(message.identity);
                    if (error)
                        return "identity." + error;
                }
                if (message.identityStatus != null && message.hasOwnProperty("identityStatus"))
                    switch (message.identityStatus) {
                    default:
                        return "identityStatus: enum value expected";
                    case 0:
                    case 1:
                    case 2:
                        break;
                    }
                return null;
            };

            /**
             * Creates a Profile message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.Profile
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.Profile} Profile
             */
            Profile.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.Profile)
                    return object;
                let message = new $root.rental.v1.Profile();
                if (object.identity != null) {
                    if (typeof object.identity !== "object")
                        throw TypeError(".rental.v1.Profile.identity: object expected");
                    message.identity = $root.rental.v1.Identity.fromObject(object.identity);
                }
                switch (object.identityStatus) {
                case "UNSUBMITTED":
                case 0:
                    message.identityStatus = 0;
                    break;
                case "PENDING":
                case 1:
                    message.identityStatus = 1;
                    break;
                case "VERIFIED":
                case 2:
                    message.identityStatus = 2;
                    break;
                }
                return message;
            };

            /**
             * Creates a plain object from a Profile message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.Profile
             * @static
             * @param {rental.v1.Profile} message Profile
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Profile.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.identity = null;
                    object.identityStatus = options.enums === String ? "UNSUBMITTED" : 0;
                }
                if (message.identity != null && message.hasOwnProperty("identity"))
                    object.identity = $root.rental.v1.Identity.toObject(message.identity, options);
                if (message.identityStatus != null && message.hasOwnProperty("identityStatus"))
                    object.identityStatus = options.enums === String ? $root.rental.v1.IdentityStatus[message.identityStatus] : message.identityStatus;
                return object;
            };

            /**
             * Converts this Profile to JSON.
             * @function toJSON
             * @memberof rental.v1.Profile
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Profile.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Profile;
        })();

        v1.Identity = (function() {

            /**
             * Properties of an Identity.
             * @memberof rental.v1
             * @interface IIdentity
             * @property {string|null} [licNumber] Identity licNumber
             * @property {string|null} [name] Identity name
             * @property {rental.v1.Gender|null} [gender] Identity gender
             * @property {number|null} [birthDateMillis] Identity birthDateMillis
             */

            /**
             * Constructs a new Identity.
             * @memberof rental.v1
             * @classdesc Represents an Identity.
             * @implements IIdentity
             * @constructor
             * @param {rental.v1.IIdentity=} [properties] Properties to set
             */
            function Identity(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Identity licNumber.
             * @member {string} licNumber
             * @memberof rental.v1.Identity
             * @instance
             */
            Identity.prototype.licNumber = "";

            /**
             * Identity name.
             * @member {string} name
             * @memberof rental.v1.Identity
             * @instance
             */
            Identity.prototype.name = "";

            /**
             * Identity gender.
             * @member {rental.v1.Gender} gender
             * @memberof rental.v1.Identity
             * @instance
             */
            Identity.prototype.gender = 0;

            /**
             * Identity birthDateMillis.
             * @member {number} birthDateMillis
             * @memberof rental.v1.Identity
             * @instance
             */
            Identity.prototype.birthDateMillis = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

            /**
             * Creates a new Identity instance using the specified properties.
             * @function create
             * @memberof rental.v1.Identity
             * @static
             * @param {rental.v1.IIdentity=} [properties] Properties to set
             * @returns {rental.v1.Identity} Identity instance
             */
            Identity.create = function create(properties) {
                return new Identity(properties);
            };

            /**
             * Encodes the specified Identity message. Does not implicitly {@link rental.v1.Identity.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.Identity
             * @static
             * @param {rental.v1.IIdentity} message Identity message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Identity.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.licNumber != null && Object.hasOwnProperty.call(message, "licNumber"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.licNumber);
                if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                if (message.gender != null && Object.hasOwnProperty.call(message, "gender"))
                    writer.uint32(/* id 3, wireType 0 =*/24).int32(message.gender);
                if (message.birthDateMillis != null && Object.hasOwnProperty.call(message, "birthDateMillis"))
                    writer.uint32(/* id 4, wireType 0 =*/32).int64(message.birthDateMillis);
                return writer;
            };

            /**
             * Encodes the specified Identity message, length delimited. Does not implicitly {@link rental.v1.Identity.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.Identity
             * @static
             * @param {rental.v1.IIdentity} message Identity message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Identity.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes an Identity message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.Identity
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.Identity} Identity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Identity.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.Identity();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.licNumber = reader.string();
                        break;
                    case 2:
                        message.name = reader.string();
                        break;
                    case 3:
                        message.gender = reader.int32();
                        break;
                    case 4:
                        message.birthDateMillis = reader.int64();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes an Identity message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.Identity
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.Identity} Identity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Identity.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies an Identity message.
             * @function verify
             * @memberof rental.v1.Identity
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Identity.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.licNumber != null && message.hasOwnProperty("licNumber"))
                    if (!$util.isString(message.licNumber))
                        return "licNumber: string expected";
                if (message.name != null && message.hasOwnProperty("name"))
                    if (!$util.isString(message.name))
                        return "name: string expected";
                if (message.gender != null && message.hasOwnProperty("gender"))
                    switch (message.gender) {
                    default:
                        return "gender: enum value expected";
                    case 0:
                    case 1:
                    case 2:
                        break;
                    }
                if (message.birthDateMillis != null && message.hasOwnProperty("birthDateMillis"))
                    if (!$util.isInteger(message.birthDateMillis) && !(message.birthDateMillis && $util.isInteger(message.birthDateMillis.low) && $util.isInteger(message.birthDateMillis.high)))
                        return "birthDateMillis: integer|Long expected";
                return null;
            };

            /**
             * Creates an Identity message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.Identity
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.Identity} Identity
             */
            Identity.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.Identity)
                    return object;
                let message = new $root.rental.v1.Identity();
                if (object.licNumber != null)
                    message.licNumber = String(object.licNumber);
                if (object.name != null)
                    message.name = String(object.name);
                switch (object.gender) {
                case "G_NOT_SPECIFIED":
                case 0:
                    message.gender = 0;
                    break;
                case "MALE":
                case 1:
                    message.gender = 1;
                    break;
                case "FEMAEL":
                case 2:
                    message.gender = 2;
                    break;
                }
                if (object.birthDateMillis != null)
                    if ($util.Long)
                        (message.birthDateMillis = $util.Long.fromValue(object.birthDateMillis)).unsigned = false;
                    else if (typeof object.birthDateMillis === "string")
                        message.birthDateMillis = parseInt(object.birthDateMillis, 10);
                    else if (typeof object.birthDateMillis === "number")
                        message.birthDateMillis = object.birthDateMillis;
                    else if (typeof object.birthDateMillis === "object")
                        message.birthDateMillis = new $util.LongBits(object.birthDateMillis.low >>> 0, object.birthDateMillis.high >>> 0).toNumber();
                return message;
            };

            /**
             * Creates a plain object from an Identity message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.Identity
             * @static
             * @param {rental.v1.Identity} message Identity
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Identity.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.licNumber = "";
                    object.name = "";
                    object.gender = options.enums === String ? "G_NOT_SPECIFIED" : 0;
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, false);
                        object.birthDateMillis = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.birthDateMillis = options.longs === String ? "0" : 0;
                }
                if (message.licNumber != null && message.hasOwnProperty("licNumber"))
                    object.licNumber = message.licNumber;
                if (message.name != null && message.hasOwnProperty("name"))
                    object.name = message.name;
                if (message.gender != null && message.hasOwnProperty("gender"))
                    object.gender = options.enums === String ? $root.rental.v1.Gender[message.gender] : message.gender;
                if (message.birthDateMillis != null && message.hasOwnProperty("birthDateMillis"))
                    if (typeof message.birthDateMillis === "number")
                        object.birthDateMillis = options.longs === String ? String(message.birthDateMillis) : message.birthDateMillis;
                    else
                        object.birthDateMillis = options.longs === String ? $util.Long.prototype.toString.call(message.birthDateMillis) : options.longs === Number ? new $util.LongBits(message.birthDateMillis.low >>> 0, message.birthDateMillis.high >>> 0).toNumber() : message.birthDateMillis;
                return object;
            };

            /**
             * Converts this Identity to JSON.
             * @function toJSON
             * @memberof rental.v1.Identity
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Identity.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Identity;
        })();

        v1.GetProfileRequest = (function() {

            /**
             * Properties of a GetProfileRequest.
             * @memberof rental.v1
             * @interface IGetProfileRequest
             */

            /**
             * Constructs a new GetProfileRequest.
             * @memberof rental.v1
             * @classdesc Represents a GetProfileRequest.
             * @implements IGetProfileRequest
             * @constructor
             * @param {rental.v1.IGetProfileRequest=} [properties] Properties to set
             */
            function GetProfileRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Creates a new GetProfileRequest instance using the specified properties.
             * @function create
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {rental.v1.IGetProfileRequest=} [properties] Properties to set
             * @returns {rental.v1.GetProfileRequest} GetProfileRequest instance
             */
            GetProfileRequest.create = function create(properties) {
                return new GetProfileRequest(properties);
            };

            /**
             * Encodes the specified GetProfileRequest message. Does not implicitly {@link rental.v1.GetProfileRequest.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {rental.v1.IGetProfileRequest} message GetProfileRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetProfileRequest.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                return writer;
            };

            /**
             * Encodes the specified GetProfileRequest message, length delimited. Does not implicitly {@link rental.v1.GetProfileRequest.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {rental.v1.IGetProfileRequest} message GetProfileRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            GetProfileRequest.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a GetProfileRequest message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.GetProfileRequest} GetProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetProfileRequest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.GetProfileRequest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a GetProfileRequest message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.GetProfileRequest} GetProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            GetProfileRequest.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a GetProfileRequest message.
             * @function verify
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            GetProfileRequest.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                return null;
            };

            /**
             * Creates a GetProfileRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetProfileRequest} GetProfileRequest
             */
            GetProfileRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetProfileRequest)
                    return object;
                return new $root.rental.v1.GetProfileRequest();
            };

            /**
             * Creates a plain object from a GetProfileRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetProfileRequest
             * @static
             * @param {rental.v1.GetProfileRequest} message GetProfileRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetProfileRequest.toObject = function toObject() {
                return {};
            };

            /**
             * Converts this GetProfileRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.GetProfileRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetProfileRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return GetProfileRequest;
        })();

        v1.ClearProfileRequest = (function() {

            /**
             * Properties of a ClearProfileRequest.
             * @memberof rental.v1
             * @interface IClearProfileRequest
             */

            /**
             * Constructs a new ClearProfileRequest.
             * @memberof rental.v1
             * @classdesc Represents a ClearProfileRequest.
             * @implements IClearProfileRequest
             * @constructor
             * @param {rental.v1.IClearProfileRequest=} [properties] Properties to set
             */
            function ClearProfileRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Creates a new ClearProfileRequest instance using the specified properties.
             * @function create
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {rental.v1.IClearProfileRequest=} [properties] Properties to set
             * @returns {rental.v1.ClearProfileRequest} ClearProfileRequest instance
             */
            ClearProfileRequest.create = function create(properties) {
                return new ClearProfileRequest(properties);
            };

            /**
             * Encodes the specified ClearProfileRequest message. Does not implicitly {@link rental.v1.ClearProfileRequest.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {rental.v1.IClearProfileRequest} message ClearProfileRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            ClearProfileRequest.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                return writer;
            };

            /**
             * Encodes the specified ClearProfileRequest message, length delimited. Does not implicitly {@link rental.v1.ClearProfileRequest.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {rental.v1.IClearProfileRequest} message ClearProfileRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            ClearProfileRequest.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a ClearProfileRequest message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.ClearProfileRequest} ClearProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            ClearProfileRequest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.ClearProfileRequest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a ClearProfileRequest message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.ClearProfileRequest} ClearProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            ClearProfileRequest.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a ClearProfileRequest message.
             * @function verify
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            ClearProfileRequest.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                return null;
            };

            /**
             * Creates a ClearProfileRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.ClearProfileRequest} ClearProfileRequest
             */
            ClearProfileRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.ClearProfileRequest)
                    return object;
                return new $root.rental.v1.ClearProfileRequest();
            };

            /**
             * Creates a plain object from a ClearProfileRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.ClearProfileRequest
             * @static
             * @param {rental.v1.ClearProfileRequest} message ClearProfileRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            ClearProfileRequest.toObject = function toObject() {
                return {};
            };

            /**
             * Converts this ClearProfileRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.ClearProfileRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            ClearProfileRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return ClearProfileRequest;
        })();

        v1.ProfileService = (function() {

            /**
             * Constructs a new ProfileService service.
             * @memberof rental.v1
             * @classdesc Represents a ProfileService
             * @extends $protobuf.rpc.Service
             * @constructor
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             */
            function ProfileService(rpcImpl, requestDelimited, responseDelimited) {
                $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
            }

            (ProfileService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = ProfileService;

            /**
             * Creates new ProfileService service using the specified rpc implementation.
             * @function create
             * @memberof rental.v1.ProfileService
             * @static
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             * @returns {ProfileService} RPC service. Useful where requests and/or responses are streamed.
             */
            ProfileService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
                return new this(rpcImpl, requestDelimited, responseDelimited);
            };

            /**
             * Callback as used by {@link rental.v1.ProfileService#getProfile}.
             * @memberof rental.v1.ProfileService
             * @typedef GetProfileCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.Profile} [response] Profile
             */

            /**
             * Calls GetProfile.
             * @function getProfile
             * @memberof rental.v1.ProfileService
             * @instance
             * @param {rental.v1.IGetProfileRequest} request GetProfileRequest message or plain object
             * @param {rental.v1.ProfileService.GetProfileCallback} callback Node-style callback called with the error, if any, and Profile
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(ProfileService.prototype.getProfile = function getProfile(request, callback) {
                return this.rpcCall(getProfile, $root.rental.v1.GetProfileRequest, $root.rental.v1.Profile, request, callback);
            }, "name", { value: "GetProfile" });

            /**
             * Calls GetProfile.
             * @function getProfile
             * @memberof rental.v1.ProfileService
             * @instance
             * @param {rental.v1.IGetProfileRequest} request GetProfileRequest message or plain object
             * @returns {Promise<rental.v1.Profile>} Promise
             * @variation 2
             */

            /**
             * Callback as used by {@link rental.v1.ProfileService#submitProfile}.
             * @memberof rental.v1.ProfileService
             * @typedef SubmitProfileCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.Profile} [response] Profile
             */

            /**
             * Calls SubmitProfile.
             * @function submitProfile
             * @memberof rental.v1.ProfileService
             * @instance
             * @param {rental.v1.IIdentity} request Identity message or plain object
             * @param {rental.v1.ProfileService.SubmitProfileCallback} callback Node-style callback called with the error, if any, and Profile
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(ProfileService.prototype.submitProfile = function submitProfile(request, callback) {
                return this.rpcCall(submitProfile, $root.rental.v1.Identity, $root.rental.v1.Profile, request, callback);
            }, "name", { value: "SubmitProfile" });

            /**
             * Calls SubmitProfile.
             * @function submitProfile
             * @memberof rental.v1.ProfileService
             * @instance
             * @param {rental.v1.IIdentity} request Identity message or plain object
             * @returns {Promise<rental.v1.Profile>} Promise
             * @variation 2
             */

            /**
             * Callback as used by {@link rental.v1.ProfileService#clearProfile}.
             * @memberof rental.v1.ProfileService
             * @typedef ClearProfileCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.Profile} [response] Profile
             */

            /**
             * Calls ClearProfile.
             * @function clearProfile
             * @memberof rental.v1.ProfileService
             * @instance
             * @param {rental.v1.IClearProfileRequest} request ClearProfileRequest message or plain object
             * @param {rental.v1.ProfileService.ClearProfileCallback} callback Node-style callback called with the error, if any, and Profile
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(ProfileService.prototype.clearProfile = function clearProfile(request, callback) {
                return this.rpcCall(clearProfile, $root.rental.v1.ClearProfileRequest, $root.rental.v1.Profile, request, callback);
            }, "name", { value: "ClearProfile" });

            /**
             * Calls ClearProfile.
             * @function clearProfile
             * @memberof rental.v1.ProfileService
             * @instance
             * @param {rental.v1.IClearProfileRequest} request ClearProfileRequest message or plain object
             * @returns {Promise<rental.v1.Profile>} Promise
             * @variation 2
             */

            return ProfileService;
        })();

        return v1;
    })();

    return rental;
})();