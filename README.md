# ATEM Switcher Go Client

Pure Go implementation of the ATEM switcher protocol by Blackmagic Design.

See the [`example`](./examples/basic/main.go) for how to use it.

_This project is in its early stages. Below is a list of commands that are currently supported._

## Supported Protocol Subset

#### Status

- ✅: Implemented
- 🤨: Behaves unexpectedly, use with caution
- 🤔: Not implemented, might be obsolete
- ⭕️: Not (yet) implemented, byte layout known
- ❓: No idea what it does or what the byte layout is like

### Switcher

| Slug   | Long name                           | Status | Notes                                                                    |
| ------ | ----------------------------------- | ------ | ------------------------------------------------------------------------ |
| `InCm` | Init complete                       | ✅     | _for internal use_                                                       |
| `_ver` | Version                             | ✅     | `state.Version`                                                          |
| `_pin` | Product Ident Name                  | ✅     | `state.Config.ProductName`                                               |
| `Warn` | Warning                             | ✅     | `state.Warning`                                                          |
| `_top` | Topology                            | 🤨     | `state.Config.Topology` - caution: contains many unknowns / wrong values |
| `_MeC` | Mix Effect Config                   | ✅     | `state.Config.MixEffect`                                                 |
| `_mpl` | Media Player Config                 | ✅     | `state.Config.MediaPlayer`                                               |
| `_MvC` | MultiView Config                    | ✅     | `state.Config.MultiViews`                                                |
| `_SSC` | SuperSource Config                  | ✅     | `state.Config.SuperSources`                                              |
| `_TlC` | TallyChannel Config                 | ✅     | `state.Config.TallyChannels`                                             |
| `_AMC` | Audio Mixer Config                  | 🤔     | seems unused?                                                            |
| `_VMC` | Video Mode Config                   | 🤔     | seems to not contain data?                                               |
| `_MAC` | Macro Config                        | ✅     | `state.Config.MacroBanks`                                                |
| `Powr` | Power                               | ✅     | `state.Power`                                                            |
| `DcOt` | DownConverter                       | ⭕️    |                                                                          |
| `VidM` | Video Mode                          | ✅     | `state.VideoMode`                                                        |
| `InPr` | Input Properties                    | ✅     | `state.Inputs` - Missing some bits                                       |
| `MvPr` | MultiViewer Properties              | ⭕️    |                                                                          |
| `MvIn` | MultiViewer Input                   | ⭕️    |                                                                          |
| `PrgI` | Program Input                       | ✅     | `state.Program`                                                          |
| `PrvI` | Preview Input                       | ✅     | `state.Preview`                                                          |
| `TrSS` | Transition Settings                 | ⭕️    |                                                                          |
| `TrPr` | Transition Preview                  | ⭕️    |                                                                          |
| `TrPs` | Transition Position                 | ⭕️    |                                                                          |
| `TMxP` | Transition Mix                      | ⭕️    |                                                                          |
| `TDpP` | Transition Dip                      | ⭕️    |                                                                          |
| `TWpP` | Transition Wipe                     | ⭕️    |                                                                          |
| `TDvP` | Transition DVE                      | ⭕️    |                                                                          |
| `TStP` | Transition Stinger                  | ⭕️    |                                                                          |
| `KeOn` | Keyer On Air                        | ⭕️    |                                                                          |
| `KeBP` | Keyer Base Properties               | ⭕️    |                                                                          |
| `KeLm` | Keyer Luma                          | ⭕️    |                                                                          |
| `KeCk` | Keyer Chroma                        | ⭕️    |                                                                          |
| `KePt` | Keyer Pattern                       | ⭕️    |                                                                          |
| `KeDV` | Keyer DVE                           | ⭕️    |                                                                          |
| `KeFS` | Keyer Fly                           | ⭕️    |                                                                          |
| `KKFP` | Keyer Fly Keyframe                  | ⭕️    |                                                                          |
| `DskB` | Downstream Keyer Base               | ⭕️    |                                                                          |
| `DskP` | Downstream Keyer Properties         | ⭕️    |                                                                          |
| `DskS` | Downstream Keyer State              | ⭕️    |                                                                          |
| `FtbP` | Fade-to-Black Properties            | ⭕️    |                                                                          |
| `FtbS` | Fade-to-Black State                 | ⭕️    |                                                                          |
| `ColV` | Color Generator                     | ⭕️    |                                                                          |
| `AuxS` | Aux Source                          | ✅     | `state.Aux`                                                              |
| `CCdo` | Camera Control Options              | ⭕️    |                                                                          |
| `CCdP` | Camera Control Properties           | ⭕️    |                                                                          |
| `CCst` | Camera Control ???                  | ❓     |                                                                          |
| `RCPS` | Clip Player State                   | ⭕️    |                                                                          |
| `MPCE` | Media Player Source                 | ✅     | `state.MediaPlayer`                                                      |
| `MPSp` | Media Pool Storage                  | ⭕️    |                                                                          |
| `MPCS` | Media Player Clip Sources           | ⭕️    |                                                                          |
| `MPAS` | Media Player Audio Sources          | ⭕️    |                                                                          |
| `MPfe` | Media Player Files                  | ✅     | `state.MediaFiles`                                                       |
| `MRPr` | Macro Run Status                    | ⭕️    |                                                                          |
| `MPrp` | Macro Properties                    | ⭕️    |                                                                          |
| `MRcS` | Macro Recording Status              | ⭕️    |                                                                          |
| `SSrc` | Super Source                        | ⭕️    |                                                                          |
| `SSBd` | Super Source Border                 | ❓     |                                                                          |
| `SSBP` | Super Source Box Properties         | ⭕️    |                                                                          |
| `AMIP` | Audio Mixer Input                   | ⭕️    |                                                                          |
| `AMMO` | Audio Mixer Master                  | ⭕️    |                                                                          |
| `AMmO` | Audio Mixer Monitor                 | ⭕️    |                                                                          |
| `AMLv` | Audio Mixer Levels                  | ⭕️    | Subscribed using `SALN`                                                  |
| `AMTl` | Audio Mixer Tally                   | ⭕️    |                                                                          |
| `AMBP` | Audio Mixer ?????                   | ❓     |                                                                          |
| `AMLP` | Audio Mixer ?????                   | ❓     |                                                                          |
| `AEBP` | Audio Effect ???                    | ❓     |
| `AIXP` | Audio Input ???                     | ❓     |
| `AICP` | Audio Input ???                     | ❓     |
| `AILP` | Audio Input ???                     | ❓     |
| `TlIn` | Tally By Index                      | ✅     | `state.TallyByIndex`                                                     |
| `TlSr` | Tally By Source                     | ✅     | `state.TallyBySource`                                                    |
| `TlFc` | Tally ?????                         | ❓     |                                                                          |
| `Time` | TimeCode Last State Change          | ✅     | `state.TimeCodeLastChange`                                               |
| `LKST` | Lock State Changed                  | ⭕️    |                                                                          |
| `RXMS` | RX Media Source???                  | ❓     |
| `RXCP` | RX ?????                            | ❓     |
| `RXSS` | RX Super Source???                  | ❓     |
| `RXCC` | RX Camera Control???                | ❓     |
| `FMTl` | ? Tally???                          | ❓     |
| `FAIP` | Fairlight Audio Input Properties??? | ❓     |
| `FASP` | Fairlight Audio ???                 | ❓     |
| `FAMP` | Fairlight Audio ???                 | ❓     |
| `FMLv` | ? Levels? Audio?                    | ❓     |
| `FDLv` | ? Levels? Audio?                    | ❓     |
| `_FAC` |                                     | ❓     |                                                                          |
| `_FEC` |                                     | ❓     |                                                                          |
| `_DVE` |                                     | ❓     |                                                                          |
| `AiVM` |                                     | ❓     |                                                                          |
| `TcLk` |                                     | ❓     |                                                                          |
| `TCCc` |                                     | ❓     |                                                                          |
| `KACk` |                                     | ❓     |                                                                          |
| `KACC` |                                     | ❓     |                                                                          |
| `CapA` |                                     | ❓     |                                                                          |
| `FMPP` |                                     | ❓     |
| `MOCP` |                                     | ❓     |

### Client
| Slug   | Long name                           | Status | Notes                                                                    |
| ------ | ----------------------------------- | ------ | ------------------------------------------------------------------------ |
| `TiRq` | Request Timecode | ✅ | `client.Timecode(...)` |
