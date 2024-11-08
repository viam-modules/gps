# [`gps` module](https://github.com/viam-modules/gps)

This [gps module](https://app.viam.com/module/viam/gps) implements the gps [rtk-serial](#Configure-your-rtk-serial-movement_sensor), [rtk-pmtk](#Configure-your-rtk-pmtk-movement_sensor), nmea, and rtk-dual-gps movement sensors using the [`rdk:component:movement_sensor` API](https://docs.viam.com/appendix/apis/components/movement_sensor/).

A global positioning system (GPS) receives signals from satellites in the earth’s orbit to determine where it is and how fast it is going.

The `rtk-serial` and [`rtk-pmtk`](https://docs.viam.com/components/movement-sensor/gps-nmea-rtk-pmtk/) movement sensor models support [NTRIP-based](https://en.wikipedia.org/wiki/Networked_Transport_of_RTCM_via_Internet_Protocol) [real time kinematic positioning (RTK)](https://en.wikipedia.org/wiki/Real-time_kinematic_positioning) GPS units ([such as these](https://www.sparkfun.com/rtk)) and RTCM versions up to 3.3.

The chip requires a correction source to get to the required positional accuracy.
The `rtk-serial` model uses an over-the-internet correction source like an RTK reference station and sends the data over a serial connection to the [board](https://docs.viam.com/components/board/).

Follow the guide to [Set up a SparkFun RTK Reference Station](https:docs.viam.com/components/movement-sensor/set-up-base-station/) to configure a SparkFun station for use with this RTK-enabled GPS movement sensor model.

## Configure your rtk-serial movement_sensor

> [!NOTE]
> Before configuring your movement_sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add movement_sensor / gps:MODEL to your machine](https://docs.viam.com/configure/#components).

On the new component panel, copy and paste the following attribute template into your movement_sensor's attributes field:

```json
{
  "serial_path": "<path_to_serial_port>",
  "serial_baud_rate": <int>,
  "ntrip_connect_attempts": <int>,
  "ntrip_mountpoint": "<identifier>",
  "ntrip_password": "<password for NTRIP server>",
  "ntrip_url": "<URL of NTRIP server>",
  "ntrip_username": "<username for NTRIP server>"
}
```

### Attributes

The following attributes are available for `viam:gps:rtk-serial` movement_sensors:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `serial_path`            | string | **Required** | The full filesystem path to the serial device, starting with <file>/dev/</file>. To find your serial device path, first connect the serial device to your machine, then:<ul><li>On Linux, run <code>ls /dev/serial/by-path/\*</code> to show connected serial devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/serial/by-path/usb-0:1.1:1.0"</code>.</li><li>On macOS, run <code>ls /dev/tty\* \| grep -i usb</code> to show connected USB serial devices, <code>ls /dev/tty\*</code> to browse all devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/ttyS0"</code>.</li></ul> |
| `serial_baud_rate`       | int    | Optional     | The rate at which data is sent from the sensor. <br> Default: `38400` |
| `ntrip_url`              | string | **Required** | The URL of the NTRIP server from which you get correction data. Connects to a base station (maintained by a third party) for RTK corrections. |
| `ntrip_username`         | string | Optional     | Username for the NTRIP server. |
| `ntrip_password`         | string | Optional     | Password for the NTRIP server. |
| `ntrip_connect_attempts` | int    | Optional     | How many times to attempt connection before timing out. <br> Default: `10` |
| `ntrip_mountpoint`       | string | Optional     | If you know of an RTK mountpoint near you, write its identifier here. It will be appended to NTRIP address string (for example, "nysnet.gov/rtcm/**NJMTPT1**") and that mountpoint's data will be used for corrections. |

## Example configuration

### `viam:gps:rtk-serial`
```json
  {
      "name": "<your-gps-rtk-serial-movement_sensor-name>",
      "model": "viam:gps:rtk-serial",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
        "serial_path": "/dev/serial/by-path/usb-0:1.1:1.0",
        "serial_baud_rate": 115200,
        "ntrip_connect_attempts": 12,
        "ntrip_mountpoint": "MNTPT",
        "ntrip_password": "pass",
        "ntrip_url": "http://ntrip/url",
        "ntrip_username": "usr"
      },
      "depends_on": []
  }
```

The `"serial_path"` filepath used in this example is specific to serial devices connected to Linux systems.
The `"serial_path"` filepath on a macOS system might resemble <file>"/dev/ttyUSB0"</file> or <file>"/dev/ttyS0"</file>.

> [!NOTE]
> How you connect your device to an NTRIP server varies by geographic region.
You will need to research the options available to you.
If you are not sure where to start, check out this [GPS-RTK2 Hookup Guide from SparkFun](https://learn.sparkfun.com/tutorials/gps-rtk2-hookup-guide/connecting-the-zed-f9p-to-a-correction-source).

## Configure your rtk-pmtk movement_sensor

> [!NOTE]
> Before configuring your movement_sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add movement_sensor / gps:MODEL to your machine](https://docs.viam.com/configure/#components).

On the new component panel, copy and paste the following attribute template into your movement_sensor's attributes field:

```json
{
  "i2c_addr": <int>,
  "i2c_baud_rate": <int>,
  "i2c_bus": "<index-of-bus-on-board>",
  "ntrip_connect_attempts": <int>,
  "ntrip_mountpoint": "<identifier>",
  "ntrip_password": "<password for NTRIP server>",
  "ntrip_url": "<URL of NTRIP server>",
  "ntrip_username": "<username for NTRIP server>"
}
```

### Attributes

The following attributes are available for `viam:gps:rtk-pmtk` movement_sensors:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `i2c_addr`               | int    | **Required** | The device's I<sup>2</sup>C address. |
| `i2c_bus`                | string | **Required** | The index of the I<sup>2</sup>C bus of the board wired to the sensor. |
| `i2c_baud_rate`          | int    | Optional     | The rate at which data is sent from the sensor. Optional. <br> Default: `38400` |
| `ntrip_url`              | string | **Required** | The URL of the NTRIP server from which you get correction data. Connects to a base station (maintained by a third party) for RTK corrections. |
| `ntrip_username`         | string | Optional     | Username for the NTRIP server. |
| `ntrip_password`         | string | Optional     | Password for the NTRIP server. |
| `ntrip_connect_attempts` | int    | Optional     | How many times to attempt connection before timing out. <br> Default: `10` |
| `ntrip_mountpoint`       | string | Optional     | If you know of an RTK mountpoint near you, write its identifier here. It will be appended to NTRIP address string (for example, "nysnet.gov/rtcm/**NJMTPT1**") and that mountpoint's data will be used for corrections. |

## Example configuration

### `viam:gps:rtk-pmtk`
```json
  {
      "name": "<your-gps-rtk-pmtk-movement_sensor-name>",
      "model": "viam:gps:rtk-pmtk",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
        "i2c_addr": 66,
        "i2c_baud_rate": 115200,
        "i2c_bus": "1",
        "ntrip_connect_attempts": 12,
        "ntrip_mountpoint": "MNTPT",
        "ntrip_password": "pass",
        "ntrip_url": "http://ntrip/url",
        "ntrip_username": "usr"
      },
      "depends_on": []
  }
```

> [!NOTE]
> How you connect your device to an NTRIP server varies by geographic region.
You will need to research the options available to you.
If you are not sure where to start, check out this [GPS-RTK2 Hookup Guide from SparkFun](https://learn.sparkfun.com/tutorials/gps-rtk2-hookup-guide/connecting-the-zed-f9p-to-a-correction-source).

## Configure your nmea movement_sensor

The `gps-nmea` movement sensor model supports [NMEA-based](https://en.wikipedia.org/wiki/NMEA_0183) GPS units and RTCM versions up to 3.3.

This GPS model uses communication standards set by the National Marine Electronics Association (NMEA).
The `gps-nmea` model can be connected using USB and send data through a serial connection to any device, or employ an I<sup>2</sup>C connection to a board.

> [!NOTE]
> Before configuring your movement_sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add movement_sensor / gps:MODEL to your machine](https://docs.viam.com/configure/#components).

On the new component panel, copy and paste the following attribute template into your movement_sensor's attributes field:

```json
{
  "connection_type": "<serial|I2C>",
  "serial_attributes": {
    "serial_path": "<your-device-path>",
    "serial_baud_rate": <int>
  },
  "i2c_attributes": {
      "i2c_bus": "<index-of-bus-on-board>",
      "i2c_addr": <int>,
      "i2c_baud_rate": <int>
  }
}
```

### Attributes

The following attributes are available for `viam:gps:nmea` movement_sensors:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `connection_type` | string  | **Required** | `"I2C"` or `"serial"`. See [Connection Attributes](#connection-attributes) below. |

### Connection attributes

You need to configure attributes to specify how the GPS connects to your computer.
You can use either serial communication (over USB) or I<sup>2</sup>C communication (through pins to a [board](https://docs.viam.com/components/board/)).

Use `connection_type` to specify `"serial"` or `"I2C"` connection in the main `attributes` config.
Then create a struct within `attributes` for either `serial_attributes` or `i2c_attributes`, respectively.

### Serial config attributes

For a movement sensor communicating over serial, you'll need to include a `serial_attributes` struct containing:

<!-- prettier-ignore -->
| Name               | Type   | Required? | Description  |
| ------------------ | ------ | --------- | ------------ |
| `serial_path` | string | **Required** | The full filesystem path to the serial device, starting with <file>/dev/</file>. To find your serial device path, first connect the serial device to your machine, then:<ul><li>On Linux, run <code>ls /dev/serial/by-path/\*</code> to show connected serial devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/serial/by-path/usb-0:1.1:1.0"</code>.</li><li>On macOS, run <code>ls /dev/tty\* \| grep -i usb</code> to show connected USB serial devices, <code>ls /dev/tty\*</code> to browse all devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/ttyS0"</code>.</li></ul> |
| `serial_baud_rate` | int    | Optional     | The rate at which data is sent from the sensor. <br> Default: `38400` |

### I2C config attributes

For a movement sensor communicating over I<sup>2</sup>C, you'll need a `i2c_attributes` struct containing:

<!-- prettier-ignore -->
| Name            | Type   | Required? | Description |
| --------------- | ------ | --------- | ----------- |
| `i2c_bus`       | string | **Required** | The index of the I<sup>2</sup>C bus on the board wired to the sensor. |
| `i2c_addr`      | int    | **Required** | The device's I<sup>2</sup>C address. |
| `i2c_baud_rate` | int    | Optional     | The rate at which data is sent from the sensor. Optional. <br> Default: `38400` |

## Example configurations

### `viam:gps:nmea` serial connection
```json
  {
      "name": "<your-gps-nmea-movement_sensor-name>",
      "model": "viam:gps:nmea",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
        "connection_type": "serial",
        "serial_attributes": {
          "serial_path": "/dev/serial/by-path/usb-0:1.1:1.0",
          "serial_baud_rate": 38400
        }
      },
      "depends_on": []
  }
```

The `"serial_path"` filepath used in this example is specific to serial devices connected to Linux systems.
The `"serial_path"` filepath on a macOS system might resemble <file>"/dev/ttyUSB0"</file> or <file>"/dev/ttyS0"</file>.

### `viam:gps:nmea` i2c connection
```json
  {
      "name": "<your-gps-nmea-movement_sensor-name>",
      "model": "viam:gps:nmea",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
        "connection_type": "I2C",
        "i2c_attributes": {
          "i2c_bus": "1",
          "i2c_addr": 111,
          "i2c_baud_rate": 38400
        }
      },
      "depends_on": []
  }
```

## Configure your rtk-dual-gps movement_sensor

The `dual-gps-rtk` model of movement sensor calculates a compass heading from two GPS movement sensors, and returns the midpoint position between the first and second GPS devices as its position.

> [!NOTE]
> Before configuring your movement_sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add movement_sensor / gps:MODEL to your machine](https://docs.viam.com/configure/#components).

On the new component panel, copy and paste the following attribute template into your movement_sensor's attributes field:

```json
{
  "first_gps": "<name-of-your-first-gps-movement-sensor>",
  "second_gps": "<name-of-your-second-gps-movement-sensor>",
  "offset_degrees": <int>
}
```

### Attributes

The following attributes are available for `viam:gps:rtk-dual-gps` movement_sensors:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `first_gps` | int | **Required** | The name you have configured for the first movement sensor you want to combine the measurements from. Must be a GPS model. |
| `second_gps` | string | **Required** | The name you have configured for the second movement sensor you want to combine the measurements from. Must be a GPS model. |
| `offset_degrees` | int | Optional | The value to offset the compass heading calculation between the two GPS devices based on their positions on the base. Calculate this as the degrees between the vector from `first_gps` to `second_gps` and the vector from the vehicle's back to the vehicle's front, counterclockwise. {{< imgproc src="/components/movement-sensor/offset_degrees.png" alt="Rand's diagram of 3 offset degree calculations." resize="600x" >}} <br> Default: `90` |

## Example configuration

### `viam:gps:rtk-dual-gps`
```json
  {
      "name": "<your-gps-rtk-dual-gps-movement_sensor-name>",
      "model": "viam:gps:rtk-dual-gps",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
        "first_gps": "gps-1",
        "second_gps": "gps-2",
        "offset_degrees": 90
      },
      "depends_on": []
  }
```

## Next Steps
- To test your movement_sensor, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your movement_sensor, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a movement_sensor component, explore [these tutorials](https://docs.viam.com/tutorials/).
