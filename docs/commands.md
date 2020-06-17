# Commands and Notes

## Sent by Switcher

| Slug   | Long name                   | State | Notes                                          |
| ------ | --------------------------- | ----- | ---------------------------------------------- |
| `InCm` | Init complete               | ✅    |                                                |
| `_ver` | Version                     | ✅    |                                                |
| `_pin` | Product Ident Name          | ✅    | null-terminted string                          |
| `Warn` | Warning                     | ✅    | null-terminted string                          |
| `_top` | Topology                    | 🤨    | caution: contains many unknowns / wrong values |
| `_MeC` | Mix Effect Config           | ✅    | # of Keyers                                    |
| `_mpl` | Media Player Config         | ✅    | Storage Capacity                               |
| `_MvC` | MultiView Config            | ✅    | # of Multiviewers                              |
| `_SSC` | SuperSource Config          | ✅    | # of Supersources                              |
| `_TlC` | TallyChannel Config         | ✅    | # of channels                                  |
| `_AMC` | Audio Mixer Config          | 🤔    | seems unused?                                  |
| `_VMC` | Video Mode Config           | 🤔    | seems to not contain data?                     |
| `_MAC` | Macro Config                | ✅    |                                                |
| `Powr` | Power                       | ✅    |                                                |
| `DcOt` | DownConverter               | ⭕️   |                                                |
| `VidM` | Video Mode                  | ✅    | see `models/videomode.go`                      |
| `InPr` | Input Properties            | ✅    | Missing some bits                              |
| `MvPr` | MultiViewer Properties      | ⭕️   |                                                |
| `MvIn` | MultiViewer Input           | ⭕️   |                                                |
| `PrgI` | Program Input               | ✅    |                                                |
| `PrvI` | Preview Input               | ✅    |                                                |
| `TrSS` | Transition Settings         | ⭕️   |                                                |
| `TrPr` | Transition Preview          | ⭕️   |                                                |
| `TrPs` | Transition Position         | ⭕️   |                                                |
| `TMxP` | Transition Mix              | ⭕️   |                                                |
| `TDpP` | Transition Dip              | ⭕️   |                                                |
| `TWpP` | Transition Wipe             | ⭕️   |                                                |
| `TDvP` | Transition DVE              | ⭕️   |                                                |
| `TStP` | Transition Stinger          | ⭕️   |                                                |
| `KeOn` | Keyer On Air                | ⭕️   |                                                |
| `KeBP` | Keyer Base Properties       | ⭕️   |                                                |
| `KeLm` | Keyer Luma                  | ⭕️   |                                                |
| `KeCk` | Keyer Chroma                | ⭕️   |                                                |
| `KePt` | Keyer Pattern               | ⭕️   |                                                |
| `KeDV` | Keyer DVE                   | ⭕️   |                                                |
| `KeFS` | Keyer Fly                   | ⭕️   |                                                |
| `KKFP` | Keyer Fly Keyframe          | ⭕️   |                                                |
| `DskB` | Downstream Keyer Base       | ⭕️   |                                                |
| `DskP` | Downstream Keyer Properties | ⭕️   |                                                |
| `DskS` | Downstream Keyer State      | ⭕️   |                                                |
| `FtbP` | Fade-to-Black Properties    | ⭕️   |                                                |
| `FtbS` | Fade-to-Black State         | ⭕️   |                                                |
| `ColV` | Color Generator             | ⭕️   |                                                |
| `AuxS` | Aux Source                  | ✅    |                                                |
| `CCdo` | Camera Control Options      | ⭕️   |                                                |
| `CCdP` | Camera Control Properties   | ⭕️   |                                                |
| `RCPS` | Clip Player State           | ⭕️   |                                                |
| `MPCE` | Media Player Source         | ✅    |                                                |
| `MPSp` | Media Pool Storage          | ⭕️   |                                                |
| `MPCS` | Media Player Clip Sources   | ⭕️   |                                                |
| `MPAS` | Media Player Audio Sources  | ⭕️   |                                                |
| `MPfe` | Media Player Files          | ✅    |                                                |
| `MRPr` | Macro Run Status            | ⭕️   |                                                |
| `MPrp` | Macro Properties            | ⭕️   |                                                |
| `MRcS` | Macro Recording Status      | ⭕️   |                                                |
| `SSrc` | Super Source                | ⭕️   |                                                |
| `SSBP` | Super Source Box Properties | ⭕️   |                                                |
| `AMIP` | Audio Mixer Input           | ⭕️   |                                                |
| `AMMO` | Audio Mixer Master          | ⭕️   |                                                |
| `AMmO` | Audio Mixer Monitor         | ⭕️   |                                                |
| `AMLv` | Audio Mixer Levels          | ⭕️   | Subscribed using `SALN`                        |
| `AMTl` | Audio Mixer Tally           | ⭕️   |                                                |
| `TlIn` | Tally By Index              |

## Sent by Client
