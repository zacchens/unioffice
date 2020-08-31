//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package zippkg ;import (_de "archive/zip";_gc "bytes";_dg "encoding/xml";_a "fmt";_bb "github.com/unidoc/unioffice";_bd "github.com/unidoc/unioffice/algo";_eg "github.com/unidoc/unioffice/common/tempstorage";_e "github.com/unidoc/unioffice/schema/soo/pkg/relationships";_b "io";_da "path";_gd "sort";_g "strings";_dc "time";);

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_de .Writer ,zipPath ,storagePath string )error {_ffc ,_bacf :=z .Create (zipPath );if _bacf !=nil {return _a .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_bacf );};_dgc ,_bacf :=_eg .Open (storagePath );if _bacf !=nil {return _a .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_bacf );};defer _dgc .Close ();_ ,_bacf =_b .Copy (_ffc ,_dgc );return _bacf ;};var _gf =[]byte {'\r','\n'};

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_ebf *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _ebf ._gg ==nil {_ebf ._gg =make (map[string ]Target );_ebf ._bg =make (map[*_e .Relationships ]string );_ebf ._f =make (map[string ]struct{});_ebf ._fe =make (map[string ]int );};_cga :=_da .Clean (filePath );if _ ,_fa :=_ebf ._f [_cga ];_fa {return false ;};_ebf ._f [_cga ]=struct{}{};_ebf ._gg [_cga ]=Target {Path :filePath ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };return true ;};

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_gg map[string ]Target ;_bg map[*_e .Relationships ]string ;_ba []Target ;_cg OnNewRelationshipFunc ;_f map[string ]struct{};_fe map[string ]int ;};func (_ef *DecodeMap )RecordIndex (path string ,idx int ){_ef ._fe [path ]=idx };

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_de .File ,dest interface{})error {_cfgb ,_ec :=f .Open ();if _ec !=nil {return _a .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_ec );};defer _cfgb .Close ();_agd :=_dg .NewDecoder (_cfgb );if _adg :=_agd .Decode (dest );_adg !=nil {return _a .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_adg );};if _gde ,_ce :=dest .(*_e .Relationships );_ce {for _fg ,_agb :=range _gde .Relationship {switch _agb .TypeAttr {case _bb .OfficeDocumentTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .OfficeDocumentType ;case _bb .StylesTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .StylesType ;case _bb .ThemeTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .ThemeType ;case _bb .SettingsTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .SettingsType ;case _bb .ImageTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .ImageType ;case _bb .CommentsTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .CommentsType ;case _bb .ThumbnailTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .ThumbnailType ;case _bb .DrawingTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .DrawingType ;case _bb .ChartTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .ChartType ;case _bb .ExtendedPropertiesTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .ExtendedPropertiesType ;case _bb .CustomXMLTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .CustomXMLType ;case _bb .WorksheetTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .WorksheetType ;case _bb .SharedStringsTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .SharedStringsType ;case _bb .TableTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .TableType ;case _bb .HeaderTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .HeaderType ;case _bb .FooterTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .FooterType ;case _bb .NumberingTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .NumberingType ;case _bb .FontTableTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .FontTableType ;case _bb .WebSettingsTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .WebSettingsType ;case _bb .FootNotesTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .FootNotesType ;case _bb .EndNotesTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .EndNotesType ;case _bb .SlideTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .SlideType ;case _bb .VMLDrawingTypeStrict :_gde .Relationship [_fg ].TypeAttr =_bb .VMLDrawingType ;};};_gd .Slice (_gde .Relationship ,func (_fbc ,_dcd int )bool {_ddg :=_gde .Relationship [_fbc ];_be :=_gde .Relationship [_dcd ];return _bd .NaturalLess (_ddg .IdAttr ,_be .IdAttr );});};return nil ;};

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_dd *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_dd ._cg =fn };

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func (_bba *DecodeMap ,_df ,_c string ,_gcd []*_de .File ,_ac *_e .Relationship ,_eb Target )error ;func MarshalXMLByType (z *_de .Writer ,dt _bb .DocType ,typ string ,v interface{})error {_gcfb :=_bb .AbsoluteFilename (dt ,typ ,0);return MarshalXML (z ,_gcfb ,v );};func (_gcf *DecodeMap )IndexFor (path string )int {return _gcf ._fe [path ]};type Target struct{Path string ;Typ string ;Ifc interface{};Index uint32 ;};const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_de .Writer ,zipPath string ,data []byte )error {_dda ,_dfd :=z .Create (zipPath );if _dfd !=nil {return _a .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_dfd );};_ ,_dfd =_b .Copy (_dda ,_gc .NewReader (data ));return _dfd ;};

// Decode loops decoding targets registered with AddTarget and calling th
func (_ed *DecodeMap )Decode (files []*_de .File )error {_ag :=1;for _ag > 0{for len (_ed ._ba )> 0{_ege :=_ed ._ba [len (_ed ._ba )-1];_ed ._ba =_ed ._ba [0:len (_ed ._ba )-1];_bac :=_ege .Ifc .(*_e .Relationships );for _ ,_cf :=range _bac .Relationship {_ddf ,_ :=_ed ._bg [_bac ];_ed ._cg (_ed ,_ddf +_cf .TargetAttr ,_cf .TypeAttr ,files ,_cf ,_ege );};};for _fb ,_ab :=range files {if _ab ==nil {continue ;};if _fcb ,_bab :=_ed ._gg [_ab .Name ];_bab {delete (_ed ._gg ,_ab .Name );if _fcd :=Decode (_ab ,_fcb .Ifc );_fcd !=nil {return _fcd ;};files [_fb ]=nil ;if _ff ,_bc :=_fcb .Ifc .(*_e .Relationships );_bc {_ed ._ba =append (_ed ._ba ,_fcb );_dec ,_ :=_da .Split (_da .Clean (_ab .Name +"\u002f\u002e\u002e\u002f"));_ed ._bg [_ff ]=_dec ;_ag ++;};};};_ag --;};return nil ;};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_de .Writer ,filename string ,v interface{})error {_gce :=&_de .FileHeader {};_gce .Method =_de .Deflate ;_gce .Name =filename ;_gce .SetModTime (_dc .Now ());_ede ,_ea :=z .CreateHeader (_gce );if _ea !=nil {return _a .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_ea );};_ ,_ea =_ede .Write ([]byte (XMLHeader ));if _ea !=nil {return _a .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_ea );};if _ea =_dg .NewEncoder (SelfClosingWriter {_ede }).Encode (v );_ea !=nil {return _a .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_ea );};_ ,_ea =_ede .Write (_gf );return _ea ;};

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_de .File ,path string )(string ,error ){_bf ,_bdg :=_eg .TempFile (path ,"\u007a\u007a");if _bdg !=nil {return "",_bdg ;};defer _bf .Close ();_gb ,_bdg :=f .Open ();if _bdg !=nil {return "",_bdg ;};defer _gb .Close ();_ ,_bdg =_b .Copy (_bf ,_gb );if _bdg !=nil {return "",_bdg ;};return _bf .Name (),nil ;};

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_cge :=_g .Split (path ,"\u002f");_gca :=_g .Join (_cge [0:len (_cge )-1],"\u002f");_gcc :=_cge [len (_cge )-1];_gca +="\u002f_\u0072\u0065\u006c\u0073\u002f";_gcc +="\u002e\u0072\u0065l\u0073";return _gca +_gcc ;};func (_cfe SelfClosingWriter )Write (b []byte )(int ,error ){_fgc :=0;_bbd :=0;for _ead :=0;_ead < len (b )-2;_ead ++{if b [_ead ]=='>'&&b [_ead +1]=='<'&&b [_ead +2]=='/'{_ge :=[]byte {};_edf :=_ead ;for _db :=_ead ;_db >=0;_db --{if b [_db ]==' '{_edf =_db ;}else if b [_db ]=='<'{_ge =b [_db +1:_edf ];break ;};};_fcf :=[]byte {};for _bcd :=_ead +3;_bcd < len (b );_bcd ++{if b [_bcd ]=='>'{_fcf =b [_ead +3:_bcd ];break ;};};if !_gc .Equal (_ge ,_fcf ){continue ;};_fff ,_cfa :=_cfe .W .Write (b [_fgc :_ead ]);if _cfa !=nil {return _bbd +_fff ,_cfa ;};_bbd +=_fff ;_ ,_cfa =_cfe .W .Write (_ca );if _cfa !=nil {return _bbd ,_cfa ;};_bbd +=3;for _fbe :=_ead +2;_fbe < len (b )&&b [_fbe ]!='>';_fbe ++{_bbd ++;_fgc =_fbe +2;_ead =_fgc ;};};};_bad ,_bde :=_cfe .W .Write (b [_fgc :]);return _bad +_bbd ,_bde ;};

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _b .Writer ;};func MarshalXMLByTypeIndex (z *_de .Writer ,dt _bb .DocType ,typ string ,idx int ,v interface{})error {_fea :=_bb .AbsoluteFilename (dt ,typ ,idx );return MarshalXML (z ,_fea ,v );};var _ca =[]byte {'/','>'};